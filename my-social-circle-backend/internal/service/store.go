package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/fabioanh/my-social-circle-backend/internal/models"
	"google.golang.org/api/iterator"
)

type Store struct {
	client *firestore.Client
}

func NewStore(projectID string) (*Store, error) {
	ctx := context.Background()
	
	// Check if emulator is set
	if emu := os.Getenv("FIRESTORE_EMULATOR_HOST"); emu != "" {
		fmt.Printf("Connecting to Firestore Emulator at %s (Project: %s)\n", emu, projectID)
	} else {
		fmt.Printf("Connecting to production Firestore (Project: %s)\n", projectID)
	}

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create firestore client: %v", err)
	}
	return &Store{client: client}, nil
}

func (s *Store) Close() {
	s.client.Close()
}

// --- Groups ---

func (s *Store) CreateGroup(ctx context.Context, userID string, group *models.Group) error {
	group.CreatedAt = time.Now()
	group.UserID = userID
	
	ref, _, err := s.client.Collection("groups").Add(ctx, group)
	if err != nil {
		return err
	}
	group.ID = ref.ID
	return nil
}

func (s *Store) UpdateGroup(ctx context.Context, groupID string, updates map[string]interface{}) error {
	_, err := s.client.Collection("groups").Doc(groupID).Update(ctx, []firestore.Update{
		{Path: "name", Value: updates["name"]},
		{Path: "description", Value: updates["description"]},
	})
	return err
}

func (s *Store) ListGroups(ctx context.Context, userID string) ([]models.Group, error) {
	iter := s.client.Collection("groups").Where("user_id", "==", userID).Documents(ctx)
	groups := []models.Group{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var g models.Group
		doc.DataTo(&g)
		g.ID = doc.Ref.ID
		groups = append(groups, g)
	}
	return groups, nil
}

func (s *Store) GetGroup(ctx context.Context, groupID string) (*models.Group, error) {
	doc, err := s.client.Collection("groups").Doc(groupID).Get(ctx)
	if err != nil {
		return nil, err
	}
	var g models.Group
	doc.DataTo(&g)
	g.ID = doc.Ref.ID
	return &g, nil
}

// --- People ---

func (s *Store) CreatePerson(ctx context.Context, userID string, person *models.Person) error {
	person.CreatedAt = time.Now()
	person.UserID = userID
	
	ref, _, err := s.client.Collection("people").Add(ctx, person)
	if err != nil {
		return err
	}
	person.ID = ref.ID
	return nil
}

func (s *Store) ListPeopleByGroup(ctx context.Context, groupID string) ([]models.Person, error) {
	// Security check: In a real app verify userID matches group owner
	iter := s.client.Collection("people").Where("group_id", "==", groupID).OrderBy("created_at", firestore.Desc).Documents(ctx)
	people := []models.Person{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var p models.Person
		doc.DataTo(&p)
		p.ID = doc.Ref.ID
		people = append(people, p)
	}
	return people, nil
}

func (s *Store) UpdatePerson(ctx context.Context, personID string, updates map[string]interface{}) error {
	_, err := s.client.Collection("people").Doc(personID).Update(ctx, []firestore.Update{
		{Path: "name", Value: updates["name"]},
	})
	return err
}

func (s *Store) DeleteFact(ctx context.Context, personID string, factID string) error {
	// First delete the fact from the subcollection
	_, err := s.client.Collection("people").Doc(personID).Collection("facts").Doc(factID).Delete(ctx)
	if err != nil {
		return err
	}

	// Update the parent person's first_fact if the deleted fact was the first one.
	// For simplicity, we'll fetch the next earliest fact and set it as first_fact.
	// If no more facts exist, set it to nil.
	iter := s.client.Collection("people").Doc(personID).Collection("facts").OrderBy("created_at", firestore.Asc).Limit(1).Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		// No more facts, set first_fact to nil
		_, err = s.client.Collection("people").Doc(personID).Update(ctx, []firestore.Update{
			{Path: "first_fact", Value: nil},
		})
		return err
	}
	if err != nil {
		return err
	}

	var f models.Fact
	doc.DataTo(&f)
	f.ID = doc.Ref.ID
	f.PersonID = personID

	_, err = s.client.Collection("people").Doc(personID).Update(ctx, []firestore.Update{
		{Path: "first_fact", Value: f},
	})
	return err
}

func (s *Store) GetPerson(ctx context.Context, personID string) (*models.Person, error) {
	doc, err := s.client.Collection("people").Doc(personID).Get(ctx)
	if err != nil {
		return nil, err
	}
	var p models.Person
	doc.DataTo(&p)
	p.ID = doc.Ref.ID
	
	// Fetch facts
	facts, err := s.ListFacts(ctx, personID)
	if err == nil {
		p.Facts = facts
	}
	
	return &p, nil
}

// --- Facts ---

func (s *Store) AddFact(ctx context.Context, fact *models.Fact) error {
	fact.CreatedAt = time.Now()
	
	// Use a transaction/batch if critical, but for now linear ops
	ref, _, err := s.client.Collection("people").Doc(fact.PersonID).Collection("facts").Add(ctx, fact)
	if err != nil {
		return err
	}
	fact.ID = ref.ID

	// Check if the person already has a first fact
	// This is a simple logic: read -> if nil -> update
	// In a high contention scenario, use a transaction.
	dsnap, err := s.client.Collection("people").Doc(fact.PersonID).Get(ctx)
	if err != nil {
		fmt.Printf("Error fetching person to check first fact: %v\n", err)
		return nil // fact added successfully anyway
	}
	var p models.Person
	dsnap.DataTo(&p)
	
	if p.FirstFact == nil {
		_, err = s.client.Collection("people").Doc(fact.PersonID).Update(ctx, []firestore.Update{
			{Path: "first_fact", Value: fact},
		})
		if err != nil {
			fmt.Printf("Failed to update first fact: %v\n", err)
		}
	}
	return nil
}

func (s *Store) ListFacts(ctx context.Context, personID string) ([]models.Fact, error) {
	iter := s.client.Collection("people").Doc(personID).Collection("facts").OrderBy("created_at", firestore.Desc).Documents(ctx)
	facts := []models.Fact{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// If subcollection doesn't exist, it might just return Done or empty. 
			// Check error types if needed.
			fmt.Printf("Error listing facts: %v\n", err)
			return nil, err
		}
		var f models.Fact
		doc.DataTo(&f)
		f.ID = doc.Ref.ID
		f.PersonID = personID // Ensure it's set
		facts = append(facts, f)
	}
	return facts, nil
}

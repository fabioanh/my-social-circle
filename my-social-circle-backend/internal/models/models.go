package models

import "time"

type Group struct {
	ID          string    `json:"id" firestore:"-"`
	Name        string    `json:"name" firestore:"name"`
	Description string    `json:"description" firestore:"description"`
	CreatedAt   time.Time `json:"created_at" firestore:"created_at"`
	UserID      string    `json:"-" firestore:"user_id"`
}

type Person struct {
	ID        string    `json:"id" firestore:"-"`
	Name      string    `json:"name" firestore:"name"`
	GroupID   string    `json:"group_id" firestore:"group_id"`
	UserID    string    `json:"-" firestore:"user_id"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
	Facts     []Fact    `json:"facts,omitempty" firestore:"-"`       // Facts fetched separately or embedded if small
	FirstFact *Fact     `json:"first_fact,omitempty" firestore:"first_fact"` // De-normalized for list views
}


type Fact struct {
	ID        string    `json:"id" firestore:"-"`
	Content   string    `json:"content" firestore:"content"`
	PersonID  string    `json:"person_id" firestore:"person_id"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
}

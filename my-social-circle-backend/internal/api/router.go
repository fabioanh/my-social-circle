package api

import (
	"encoding/json"
	"net/http"

	"github.com/fabioanh/my-social-circle-backend/internal/models"
	"github.com/fabioanh/my-social-circle-backend/internal/service"
)

type Router struct {
	store *service.Store
}

func NewRouter(store *service.Store) http.Handler {
	r := &Router{store: store}
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /health", r.HealthCheck)
	
	// Groups
	mux.HandleFunc("POST /groups", r.CreateGroup)
	mux.HandleFunc("GET /groups", r.ListGroups)
	mux.HandleFunc("GET /groups/{id}", r.GetGroup)
	mux.HandleFunc("PUT /groups/{id}", r.UpdateGroup)

	// People
	mux.HandleFunc("POST /people", r.CreatePerson)
	mux.HandleFunc("GET /groups/{id}/people", r.ListPeopleByGroup)
	mux.HandleFunc("GET /people/{id}", r.GetPerson)
	mux.HandleFunc("PUT /people/{id}", r.UpdatePerson)

	// Facts
	mux.HandleFunc("POST /people/{id}/facts", r.AddFact)
	mux.HandleFunc("DELETE /people/{id}/facts/{factId}", r.DeleteFact)

	return CorsMiddleware(mux)
}

func (r *Router) HealthCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (r *Router) CreateGroup(w http.ResponseWriter, req *http.Request) {
	var g models.Group
	if err := json.NewDecoder(req.Body).Decode(&g); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Hardcoded user ID until auth is implemented
	userID := "demo-user"
	
	if err := r.store.CreateGroup(req.Context(), userID, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)
}

func (r *Router) ListGroups(w http.ResponseWriter, req *http.Request) {
	userID := "demo-user"
	groups, err := r.store.ListGroups(req.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func (r *Router) UpdateGroup(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	var updates map[string]interface{}
	if err := json.NewDecoder(req.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := r.store.UpdateGroup(req.Context(), id, updates); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (r *Router) GetGroup(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	g, err := r.store.GetGroup(req.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)
}

func (r *Router) CreatePerson(w http.ResponseWriter, req *http.Request) {
	var p models.Person
	if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID := "demo-user"
	
	if err := r.store.CreatePerson(req.Context(), userID, &p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (r *Router) ListPeopleByGroup(w http.ResponseWriter, req *http.Request) {
	groupID := req.PathValue("id")
	people, err := r.store.ListPeopleByGroup(req.Context(), groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func (r *Router) GetPerson(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	p, err := r.store.GetPerson(req.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (r *Router) UpdatePerson(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	var updates map[string]interface{}
	if err := json.NewDecoder(req.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := r.store.UpdatePerson(req.Context(), id, updates); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (r *Router) AddFact(w http.ResponseWriter, req *http.Request) {
	personID := req.PathValue("id")
	var f models.Fact
	if err := json.NewDecoder(req.Body).Decode(&f); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	f.PersonID = personID
	
	if err := r.store.AddFact(req.Context(), &f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(f)
}

func (r *Router) DeleteFact(w http.ResponseWriter, req *http.Request) {
	personID := req.PathValue("id")
	factID := req.PathValue("factId")

	if err := r.store.DeleteFact(req.Context(), personID, factID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Enable CORS middleware


func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/AbhiramiRajeev/task-orchestrator/auth-service/internal/model"
	bcrpypt "golang.org/x/crypto/bcrypt"
)

var (
	mu     sync.Mutex
	users  = make([]model.User, 0)
	nextID = 1
)

func Registerhandler(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	hash, err := bcrpypt.GenerateFromPassword([]byte(req.Password), bcrpypt.DefaultCost)
	if err != nil {
		http.Error(w, error.Error(err), http.StatusBadRequest)
	}
	mu.Lock()
	user := model.User{
		Username: req.Username,
		Password: string(hash),
	}
	users = append(users, user)
	nextID++
	mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

func Loginhandler(w http.ResponseWriter, r http.Response) {
	var req model.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	mu.Lock()
	var user *model.User

	for _, u := range users {
		if u.Username == req.Username {
			user = &u
			break
		}
	}
	mu.Unlock()
	if user == nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	if err := bcrpypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})

}

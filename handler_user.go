package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/logeshwhatnot/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	parameters := params{}
	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error decoding the body: %v", err))
		return
	}

	user, err := apiCfg.db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.NullUUID{UUID: uuid.New(), Valid: true},
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      parameters.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return

	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}

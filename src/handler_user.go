package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rssagregate/src/internal/database"
	"time"

	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	par := params{}
	err := decoder.Decode(&par)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't decode to user params %v", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      par.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't create user %v", err))
		return
	}

	respondWithJson(w, 200, databaseUsertoUser(user))
}

func (apiConfig *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, 200, databaseUsertoUser(user))
}

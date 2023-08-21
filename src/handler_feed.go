package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rssagregate/src/internal/database"
	"time"

	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	par := params{}
	err := decoder.Decode(&par)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't decode to feed params %v", err))
		return
	}

	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now() ,
		Name:      par.Name,
		Url:       par.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't create feed %v", err))
		return
	}

	respondWithJson(w, 200, databaseFeedtoFeed(feed))
}

func (apiConfig *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiConfig.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't find any feed %v", err))
		return
	}

	respondWithJson(w, 200, databaseFeedstoFeeds(feeds))
}

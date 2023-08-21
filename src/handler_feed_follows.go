package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rssagregate/src/internal/database"
	"time"

	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	par := params{}
	err := decoder.Decode(&par)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't decode to feed_follows params %v", err))
		return
	}

	feed_follow, err := apiConfig.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    par.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't create feed follows %v", err))
		return
	}

	respondWithJson(w, 200, databaseFeedFollowtoFeedFollow(feed_follow))
}

func (apiConfig *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follows, err := apiConfig.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't find any feed follows %v", err))
		return
	}

	respondWithJson(w, 200, databaseFeedFollowsToFeedFollows(feed_follows))
}

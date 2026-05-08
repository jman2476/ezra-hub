package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jman2476/ezra-hub/app/server/internal/auth"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
)

func (cfg *apiConfig) handlerNewEvent(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Name      string    `json:"name"`
		Type      string    `json:"type"`
		OccursOn  time.Time `json:"occurs_on"`
		ExpiresAt time.Time `json:"expires_at"`
	}

	log.Printf("POST /api/events")

	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Forbidden", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "JWT Expired", err)
		return
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding event body", err)
		return
	}

	eventArgs := database.CreateEventParams{
		Name:      params.Name,
		OwnerID:   userID,
		Type:      params.Type,
		OccursOn:  params.OccursOn,
		ExpiresAt: params.ExpiresAt,
	}

	event, err := cfg.db.CreateEvent(req.Context(), eventArgs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Database error: unable to create event", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, mapEvent(event))
}

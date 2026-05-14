package main

import (
	"database/sql"
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
		Category  string    `json:"category"`
		OccursOn  time.Time `json:"occurs_on"`
		ExpiresAt time.Time `json:"expires_at"`
		MinVol    int32     `json:"min_volunteer"`
		MaxVol    int32     `json:"max_volunteer"`
		Desc      string    `json:"description"`
	}

	log.Printf("POST /api/events")

	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Forbidden", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "JWT Expired", err)
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
		Category:  database.Genre(validateGenre(params.Category)),
		OccursOn:  params.OccursOn,
		ExpiresAt: params.ExpiresAt,
		MinVolunteers: sql.NullInt32{
			Int32: params.MinVol,
			Valid: params.MinVol != 0},
		MaxVolunteers: sql.NullInt32{
			Int32: params.MaxVol,
			Valid: params.MaxVol != 0},
		Description: params.Desc,
	}

	event, err := cfg.db.CreateEvent(req.Context(), eventArgs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Database error: unable to create event", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, mapEvent(event))
}

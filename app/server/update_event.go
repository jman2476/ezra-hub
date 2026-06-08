package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
	"github.com/jman2476/ezra-hub/app/server/internal/msgbroker"
	"github.com/jman2476/ezra-hub/pkg/routing"
)

func (cfg *apiConfig) handlerUpdateEvent(w http.ResponseWriter, req *http.Request, userID uuid.UUID) {
	log.Println("PUT /api/events/{id}")

	type parameters struct {
		Name      string    `json:"name"`
		Category  string    `json:"category"`
		OccursOn  time.Time `json:"occurs_on"`
		ExpiresAt time.Time `json:"expires_at"`
		MinVol    int32     `json:"min_volunteer"`
		MaxVol    int32     `json:"max_volunteer"`
		Desc      string    `json:"description"`
		OldType   string    `json:"old_type"`
		Location  string    `json:"location"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error decoding request body", err)
		return
	}

	eventID := req.PathValue("id")
	eventUUID, err := uuid.Parse(eventID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid event identifies", err)
		return
	}

	eventArgs := database.UpdateEventByIDParams{
		Name:        params.Name,
		Description: params.Desc,
		Category:    database.Genre(validateGenre(params.Category)),
		OccursOn:    params.OccursOn,
		ExpiresAt:   params.ExpiresAt,
		MinVolunteers: sql.NullInt32{
			Int32: params.MinVol,
			Valid: params.MinVol != 0},
		MaxVolunteers: sql.NullInt32{
			Int32: params.MaxVol,
			Valid: params.MaxVol != 0},
		ID:       eventUUID,
		Location: params.Location,
	}

	event, err := cfg.db.UpdateEventByID(req.Context(), eventArgs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Database error: unable to update event", err)
		return
	}

	creator_name, err := cfg.db.GetUserNameOnly(req.Context(), event.OwnerID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Database error: cannot get creator's name. Event still created", err)
		return
	}

	err = msgbroker.PublishJSON(
		cfg.channel,
		routing.ExchangeEzraTopic,
		string(event.Category)+".update."+creator_name,
		mapEvent(event), cfg.db,
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Message broker error: unable to publish updated event", err)
		return
	}
	if params.OldType != string(event.Category) {
		err = msgbroker.PublishJSON(
			cfg.channel,
			routing.ExchangeEzraTopic,
			string(params.OldType)+".update."+creator_name,
			mapEvent(event), cfg.db,
		)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Message broker error: unable to publish updated event on old event queue", err)
			return
		}
	}

	respondWithJSON(w, http.StatusOK, mapEvent(event))
}

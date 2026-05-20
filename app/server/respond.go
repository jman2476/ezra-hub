package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
	"github.com/jman2476/ezra-hub/app/server/internal/msgbroker"
	"github.com/jman2476/ezra-hub/pkg/routing"
)

func (cfg *apiConfig) handlerRespondEvent(w http.ResponseWriter, req *http.Request, userID uuid.UUID) {
	log.Println("PATCH /api/events/{id}")

	type parameters struct {
		Available bool `json:"available"`
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
		respondWithError(w, http.StatusBadRequest, "Invalid event identifier", err)
	}

	var event database.Event

	if params.Available {
		responderParams := database.AddEventResponderParams{
			ArrayAppend: userID,
			ID:          eventUUID,
		}

		event, err = cfg.db.AddEventResponder(req.Context(), responderParams)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Database error updating event respondants", err)
			return
		}

	} else {
		responderParams := database.RemoveEventResponderParams{
			ArrayRemove: userID,
			ID:          eventUUID,
		}
		event, err = cfg.db.RemoveEventResponder(req.Context(), responderParams)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Database error updating event respondants", err)
		}
	}

	creator_name, err := cfg.db.GetUserNameOnly(req.Context(), event.OwnerID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Database error: cannot get creator's name. Event still updated with responder", err)
		return
	}
	log.Printf("\r%s created %s\n", creator_name, event.Name)

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

	respondWithJSON(w, http.StatusNoContent, struct{}{})
}

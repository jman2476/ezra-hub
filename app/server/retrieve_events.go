package main

import (
	"log"
	"net/http"
	"slices"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetEventsByType(w http.ResponseWriter, req *http.Request, userID uuid.UUID) {
	log.Printf("GET /api/events?type=")

	var genres []Genre

	types := req.URL.Query()["type"]
	if len(types) == 0 {
		respondWithJSON(w, http.StatusNoContent, struct{}{})
	}

	if slices.Contains(types, "all") {
		genres = []Genre{
			GenreRide, GenreShopping, GenreCheckIn, GenreMeal, GenreGathering, GenreOther,
		}
	} else {
		for _, t := range types {
			genres = append(genres, Genre(t))
		}
	}

	events, err := cfg.db.GetEventsByCategory(req.Context(), mapGenres2DB(genres))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error retrieving events from database", err)
		return
	}

	respondWithJSON(w, http.StatusOK, mapEvents(events))
}

func (cfg *apiConfig) handlerGetEventsbyUser(w http.ResponseWriter, req *http.Request, userID uuid.UUID) {
	log.Printf("GET /api/events/users/")

	events, err := cfg.db.GetEventsByOwner(req.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting user's events from database", err)
		return
	}

	if len(events) == 0 {
		respondWithJSON(w, http.StatusNoContent, struct{}{})
		return
	}

	respondWithJSON(w, http.StatusOK, mapEvents(events))
}

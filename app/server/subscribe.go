package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"

	"github.com/google/uuid"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
)

var subTypes = []string{
	"ride", "shopping", "check-in", "meal", "gathering", "other",
}

func (cfg *apiConfig) handlerSubscribe(w http.ResponseWriter, req *http.Request, userID uuid.UUID) {
	log.Println("PATCH /api/users")

	type parameters struct {
		Queues map[string]int `json:"subscriptions"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error decoding request body", err)
		return
	}

	currentSubs, err := cfg.db.GetUserSubsbyID(req.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Can't get user's subscriptions", err)
		return
	}
	log.Printf("Paramters struct %v", params)
	log.Printf("Params queues: %v", params.Queues)
	subsToAdd := getNewSubscriptions(params.Queues, currentSubs.Subs)

	subParams := database.SetSubscriptionbyIDParams{
		ID:   userID,
		Subs: slices.Concat(currentSubs.Subs, subsToAdd),
	}

	err = cfg.db.SetSubscriptionbyID(req.Context(), subParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to set new subscriptions", err)
		return
	}

	respondWithJSON(w, http.StatusOK, subsToAdd)
}

func getNewSubscriptions(new map[string]int, old []database.Subscription) []database.Subscription {
	var newList []database.Subscription
	for _, sub := range old {
		if _, ok := new[string(sub)]; ok {
			new[string(sub)] -= 1
		}
	}
	for key, val := range new {
		if val == 1 && slices.Contains(subTypes, key) {
			newList = append(newList, database.Subscription(key))
		}
	}
	log.Printf("new: %v\r\nold: %v\r\nnewList: %v", new, old, newList)

	return newList
}

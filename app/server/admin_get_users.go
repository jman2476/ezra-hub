package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerGetUsers(w http.ResponseWriter, req *http.Request) {
	log.Println("GET /admin/users")

	users, err := cfg.db.GetUsers(req.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Database error", err)
		return
	}

	mappedUsers := []User{}
	for _, u := range users {
		mappedUsers = append(mappedUsers, mapUser(u, "", ""))
	}

	respondWithJSON(w, http.StatusOK, mappedUsers)
}

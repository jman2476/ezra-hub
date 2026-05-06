package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jman2476/ezra-hub/app/server/internal/database"
)

func (cfg *apiConfig) handerLogIn(w http.ResponseWriter, req *http.Request) {
	log.Println("POST /api/login")

	type parameters struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Log in error: decoding login data", err)
		return
	}

	userArgs := database.GetUserByNameEmailParams{
		Name:  params.Name,
		Email: params.Email,
	}
	user, err := cfg.db.GetUserByNameEmail(req.Context(), userArgs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Database retreival error", err)
		return
	}

	respondWithJSON(w, http.StatusOK, mapUser(user))
}

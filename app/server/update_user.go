package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
)

func (cfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, req *http.Request, userID uuid.UUID) {
	log.Println("PATCH /api/users")

	type parameters struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error decoding request body", err)
		return
	}

	phoneNumber, ok := validatePhoneNumber(w, params.PhoneNumber)
	if !ok {
		// Already responded w/ error
		return
	}

	email, ok := validateEmail(w, params.Email)
	if !ok {
		// Already responded w/ error
		return
	}

	updateParams := database.UpdateUserbyIDParams{
		Name:        params.Name,
		Email:       email,
		PhoneNumber: phoneNumber,
		ID:          userID,
	}

	user, err := cfg.db.UpdateUserbyID(req.Context(), updateParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error updating user in database", err)
		return
	}

	respondWithJSON(w, http.StatusOK, mapUser(user, "", ""))
}

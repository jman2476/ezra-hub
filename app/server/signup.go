package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"

	"github.com/jman2476/ezra-hub/internal/database"
	"github.com/ttacon/libphonenumber"
)

func (cfg *apiConfig) handlerNewUser(w http.ResponseWriter, req *http.Request) {
	log.Printf("POST /api/users")

	type parameters struct {
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
		Email       string `json:"email"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding user sign up data", err)
		return
	}

	phoneNumber, err := libphonenumber.Parse(params.PhoneNumber, "US")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid phone number", err)
		return
	}
	phoneStr := fmt.Sprintf("+%d %d", phoneNumber.CountryCode, phoneNumber.NationalNumber)

	email, err := mail.ParseAddress(params.Email)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid email address", err)
		return
	}

	userArgs := database.CreateUserParams{
		Name:        params.Name,
		PhoneNumber: phoneStr,
		Email:       email.Address,
	}

	user, err := cfg.db.CreateUser(req.Context(), userArgs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating user", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

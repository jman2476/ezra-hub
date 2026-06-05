package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/jman2476/ezra-hub/app/server/internal/auth"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
	"github.com/nyaruka/phonenumbers"
)

func (cfg *apiConfig) handlerNewUser(w http.ResponseWriter, req *http.Request) {
	log.Printf("POST /api/users")

	type parameters struct {
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		Address     string `json:"address"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding user sign up data", err)
		return
	}

	if params.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Missing password for new user", err)
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

	hash, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Auth hashing error", err)
		return
	}

	userArgs := database.CreateUserParams{
		Name:           params.Name,
		PhoneNumber:    phoneNumber,
		Email:          email,
		HashedPassword: hash,
		Address:        params.Address,
	}

	user, err := cfg.db.CreateUser(req.Context(), userArgs)
	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint \"users_name_email_key\"") {
			respondWithError(w, http.StatusBadRequest, "Name&email combination in use by another user", err)
			return
		}
		if strings.Contains(err.Error(), "violates unique constraint \"users_name_phone_number_key\"") {
			respondWithError(w, http.StatusBadRequest, "Name&phonenumber combination in use by another user", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Error creating user", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, mapUser(user, "", ""))
}

func validatePhoneNumber(w http.ResponseWriter, pn string) (string, bool) {
	phoneNumber, err := phonenumbers.Parse(pn, "US")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid phone number", err)
		return "", false
	}

	return fmt.Sprintf("+%d %d", *phoneNumber.CountryCode, *phoneNumber.NationalNumber), true
}

func validateEmail(w http.ResponseWriter, em string) (string, bool) {
	email, err := mail.ParseAddress(em)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid email address", err)
		return "", false
	}
	return email.Address, true
}

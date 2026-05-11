package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jman2476/ezra-hub/app/server/internal/auth"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
)

func (cfg *apiConfig) handerLogIn(w http.ResponseWriter, req *http.Request) {
	log.Println("POST /api/login")

	type parameters struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Log in error: decoding login data", err)
		return
	}

	userArgs := database.GetUserforLoginParams{
		Name:  params.Name,
		Email: params.Email,
	}
	user, err := cfg.db.GetUserforLogin(req.Context(), userArgs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Database retreival error", err)
		return
	}

	valid, err := auth.CheckPassword(params.Password, user.HashedPassword)
	if !valid || err != nil {
		respondWithError(w, http.StatusForbidden, "Invalid login credentials", err)
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.secret, time.Minute)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error making authentication token", err)
		return
	}

	refreshArgs := database.CreateRefreshTokenParams{
		Token:  auth.MakeRefreshToken(),
		UserID: user.ID,
	}

	refresh, err := cfg.db.CreateRefreshToken(req.Context(), refreshArgs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating refresh token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, mapUser(user, token, refresh.Token))
}

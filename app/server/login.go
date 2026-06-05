package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"time"

	"github.com/jman2476/ezra-hub/app/server/internal/auth"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
	"github.com/jman2476/ezra-hub/app/server/internal/msgbroker"
	"github.com/jman2476/ezra-hub/pkg/routing"
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
	log.Println("decoding")
	userArgs := database.GetUserforLoginParams{
		Name:  params.Name,
		Email: params.Email,
	}
	user, err := cfg.db.GetUserforLogin(req.Context(), userArgs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Database retreival error", err)
		return
	}
	log.Println("Got user")
	valid, err := auth.CheckPassword(params.Password, user.HashedPassword)
	if !valid || err != nil {
		respondWithError(w, http.StatusForbidden, "Invalid login credentials", err)
		return
	}
	log.Println("Got password")
	var token string
	if slices.Contains(cfg.args, "shortJWT") {
		log.Println("Creating short JWT: 3 minute lifetime")
		token, err = auth.MakeJWT(user.ID, cfg.secret, time.Minute*3)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error making authentication token", err)
			return
		}
	} else {
		token, err = auth.MakeJWT(user.ID, cfg.secret, time.Hour)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error making authentication token", err)
			return
		}

	}
	log.Println("verified JWT")
	refreshArgs := database.CreateRefreshTokenParams{
		Token:  auth.MakeRefreshToken(),
		UserID: user.ID,
	}

	refresh, err := cfg.db.CreateRefreshToken(req.Context(), refreshArgs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating refresh token", err)
		return
	}
	log.Println("created refresh token")
	err = msgbroker.PublishJSON(
		cfg.channel,
		routing.ExchangeEzraDirect,
		routing.ActiveUserKey,
		routing.ActiveUser{Name: user.Name},
		cfg.db,
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Message broker error: unable to publish user login", err)
		return
	}
	log.Println("published")
	respondWithJSON(w, http.StatusOK, mapUser(user, token, refresh.Token))
}

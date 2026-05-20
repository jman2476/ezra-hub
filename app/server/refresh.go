package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"slices"
	"time"

	"github.com/jman2476/ezra-hub/app/server/internal/auth"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
)

var (
	errExpiredRefreshToken = errors.New("refresh token: token expired")
	errRevokedRefreshToken = errors.New("refresh token: token revoked")
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, req *http.Request) {
	type response struct {
		Token string `json:"jwt_token"`
	}

	log.Printf("POST /api/refresh")

	refreshToken, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Forbidden: Bad refresh token", err)
		return
	}

	refreshData, err := cfg.db.GetRefreshToken(req.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Forbidden: Invalid refresh token", err)
		return
	}

	if valid, err := cfg.verifyRefreshToken(refreshData); !valid {
		respondWithError(w, http.StatusUnauthorized, "Forbidden: Expired token", err)
		return
	}

	var newToken string
	if slices.Contains(cfg.args, "shortJWT") {
		log.Println("Refreshing short JWT: 3 minute lifetime")
		newToken, err = auth.MakeJWT(refreshData.UserID, cfg.secret, time.Minute*3)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Forbidden", err)
			return
		}
	} else {
		newToken, err = auth.MakeJWT(refreshData.UserID, cfg.secret, time.Hour)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Forbidden", err)
			return
		}
	}

	respondWithJSON(w, http.StatusOK, response{Token: newToken})
}

func (cfg *apiConfig) verifyRefreshToken(token database.Refreshtoken) (bool, error) {
	if token.RevokedAt.Valid {
		return false, errRevokedRefreshToken
	}

	timeRemaining := time.Until(token.ExpiresAt)
	if timeRemaining <= 0 {
		cfg.db.RevokeToken(context.Background(), token.Token)
		return false, errExpiredRefreshToken
	}

	return true, nil
}

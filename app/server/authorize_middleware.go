package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jman2476/ezra-hub/app/server/internal/auth"
)

func (cfg *apiConfig) authorize(handler func(http.ResponseWriter, *http.Request, uuid.UUID)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		token, err := auth.GetBearerToken(req.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Forbidden", err)
			return
		}

		userID, err := auth.ValidateJWT(token, cfg.secret)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "JWT Expired", err)
			return
		}

		handler(w, req, userID)

	}

}

package main

import (
	"errors"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerResetDB(w http.ResponseWriter, req *http.Request) {
	log.Println("POST /admin/reset")
	if cfg.platform != "dev" {
		log.Println("Access denied")
		respondWithError(w, http.StatusForbidden, "I'm sorry Dave, I cannot let you do that", errors.New("Must be in 'dev' environment to clear users table"))
		return
	}

	err := cfg.db.ClearUsers(req.Context())
	if err != nil {
		log.Println("Reset users table error")
		respondWithError(w, http.StatusInternalServerError, "Error clearing users table", err)
		return
	}
	log.Println("Reset users table successful")
	response := struct {
		Msg string `json:"msg"`
	}{
		Msg: "Users table successfully reset",
	}

	respondWithJSON(w, http.StatusOK, response)
}

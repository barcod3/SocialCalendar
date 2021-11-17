package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ErrorMessage struct {
	Message string
	Error   error
	Code    int
}

func JSONError(w http.ResponseWriter, e ErrorMessage) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(e.Code)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		logrus.WithError(err).Error("Couldnt write json error message")
	}
}

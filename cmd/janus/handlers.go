package main

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/hellofresh/janus/pkg/config"
	"github.com/hellofresh/janus/pkg/errors"
	"github.com/hellofresh/janus/pkg/response"
)

func Home(app config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, fmt.Sprintf("Welcome to %s, this is version %s", app.Name, app.Version))
	}
}

// RecoveryHandler handler for the apis
func RecoveryHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	switch internalErr := err.(type) {
	case *errors.Error:
		log.WithFields(log.Fields{"code": internalErr.Code, "error": internalErr.Error()}).
			Warning("Internal error hadled")
		response.JSON(w, internalErr.Code, internalErr.Error())
	default:
		log.WithField("error", err).Error("Internal server error handled")
		response.JSON(w, http.StatusInternalServerError, err)
	}
}

func Heartbeat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, nil)
	}
}

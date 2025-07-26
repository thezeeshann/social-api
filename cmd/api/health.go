package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":  "ok",
		"env":     app.config.addr,
		"version": version,
	}
	if err := writeJson(w, http.StatusOK, data); err != nil {
		writeJsonError(w, http.StatusInternalServerError, "err.Error()")
	}

}

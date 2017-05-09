package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	result := map[string]string{
		"status": "ok",
	}

	renderResponse(result, w)
}

func renderResponse(res interface{}, w http.ResponseWriter) {
	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
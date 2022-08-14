package api

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(Response[string]{Status: "error", Content: err.Error()})
}

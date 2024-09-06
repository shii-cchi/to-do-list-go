package delivery

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondWithJSON sends a JSON response with the given HTTP status code and payload to the client.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf(ErrMarshalingJSON+": %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

// RespondWithError sends a JSON response with an error message and status code to the client.
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

// This script contains the function to return a JSON response

package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, body any) {
	// Set Header
	w.Header().Set("Content-Type", "application/json")
	// Set Status
	w.WriteHeader(status)

	if body == nil {
		w.WriteHeader(status)
		return
	}

	// marshal body
	bytes, err := json.Marshal(body)
	if err != nil {
		// default error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bytes)

}

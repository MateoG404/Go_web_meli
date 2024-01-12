// This script contains the response struct for the text handler

package response

import "net/http"

// Text writes text response
func TextResponse(w http.ResponseWriter, code int, body string) {
	// set header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// set status code
	w.WriteHeader(code)

	// write body
	w.Write([]byte(body))
}

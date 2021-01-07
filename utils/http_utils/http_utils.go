package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/a-soliman/bookstore_utils-go/rest_errors"
)

// ResponseJSON a util function to set the appropriate header and encode the response body
func ResponseJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

// ResponseError a util function to set the appropriate header and encode the response body
func ResponseError(w http.ResponseWriter, err rest_errors.RestErr) {
	ResponseJSON(w, err.Status(), err)
}

package helpers

// Todo-> Refine response wrapper
import (
	"encoding/json"
	"net/http"
)

// Main wrapper for response
//
// swagger:response ResponseWrapper
type ResponseWrapper struct {
	// in:body
	Body Response
}

// Wrapper for all responses form the system
//
type Response struct {
	// Status of the request
	Status string `json:"status"`
	// Description on what happened to the request
	Message string `json:"message"`
	// Data returned from API as part of request
	Model interface{} `json:"model"`
}

func (r *Response) ToResponse(writer http.ResponseWriter) {
	e := json.NewEncoder(writer)
	err := e.Encode(r)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}

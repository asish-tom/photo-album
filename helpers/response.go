package helpers

import (
	"encoding/json"
	"net/http"
)

// swagger:response ResponseModel
// Wrapper for all responses form the system
type ResponseModel struct {
	// in: body
	Status  string
	Message string
	Model   interface{}
}

func (r *ResponseModel) ToResponse(writer http.ResponseWriter) {
	e := json.NewEncoder(writer)
	err := e.Encode(r)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}

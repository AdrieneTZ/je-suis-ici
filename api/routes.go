package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

// SetupRoutes configures all API routes
func SetupRoutes(handler *Handler) *mux.Router {
	r := mux.NewRouter()

	// API version prefix
	api := r.PathPrefix("/api/v1").Subrouter()

	// Check-in routes
	api.HandleFunc("/checkins", handler.CreateCheckIn).Methods(http.MethodPost)

	// Serve the Lexicon schemas
	api.HandleFunc("/xrpc/app.checkin.post", HandleCheckInSchema).Methods(http.MethodGet)

	return r
}

// Handle check-in schema request
func HandleCheckInSchema(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"lexicon": 1, "id": "app.checkin.post", "type": "record", "description": "user's check-in record", "record": {"type": "object", "required": ["userDid", "location", "createdAt"], "properties": {"id": {"type": "string"}, "userDid": {"type": "string", "format": "did"}, "text": {"type": "string", "maxLength": 5000}, "location": {"type": "object", "required": ["latitude", "longitude", "locationName"], "properties": {"latitude": {"type": "number"}, "longitude": {"type": "number"}, "locationName": {"type": "string"}}}, "images": {"type": "array", "items": {"type": "string", "format": "uri"}}, "createdAt": {"type": "string", "format": "datetime"}, "updatedAt": {"type": "string", "format": "datetime"}}}}`))
}

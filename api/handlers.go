package api

import (
	"encoding/json"
	"je-suis-ici/models"
	"je-suis-ici/services"
	"net/http"
)

type Handler struct {
	checkInService *services.CheckInService
}

func NewHandler(checkInService *services.CheckInService) *Handler {
	return &Handler{
		checkInService: checkInService,
	}
}

// CreateCheckIn handles the creation of a new check-in
func (h *Handler) CreateCheckIn(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID       int      `json:"userId"`
		UserDID      string   `json:"userDid"`
		Text         string   `json:"text"`
		Latitude     float64  `json:"latitude"`
		Longitude    float64  `json:"longitude"`
		LocationName string   `json:"locationName"`
		Images       []string `json:"images"`
	}

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.UserID <= 0 {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}
	if req.UserDID == "" {
		http.Error(w, "userDid is required", http.StatusBadRequest)
		return
	}
	if req.LocationName == "" {
		http.Error(w, "locationName is required", http.StatusBadRequest)
		return
	}

	// Create CheckIn object
	checkIn := &models.CheckIn{
		UserID:       req.UserID,
		UserDID:      req.UserDID,
		Text:         req.Text,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		LocationName: req.LocationName,
		Images:       req.Images,
	}

	// Call service to create check-in
	id, err := h.checkInService.CreateCheckIn(checkIn)
	if err != nil {
		http.Error(w, "Failed to create check-in: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: successfully create a checkin, but need to fix retrieve fail error
	// Get created check-in
	createdCheckIn, err := h.checkInService.GetCheckIn(id)
	if err != nil {
		http.Error(w, "Failed to retrieve created check-in", http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCheckIn)
}

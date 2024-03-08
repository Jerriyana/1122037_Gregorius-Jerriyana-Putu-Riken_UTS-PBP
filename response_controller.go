package controllers

import (
	m "UTS/models"
	"encoding/json"
	"net/http"
)

// Send Error
func sendErrorResponse(w http.ResponseWriter, status int, message string) {
	var response m.ErrorResponse
	response.Status = status
	response.Message = message

	// Mengirimkan response JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// Send Success
func sendSuccessResponse(w http.ResponseWriter, status int, message string) {
	var response m.SuccessResponse
	response.Status = status
	response.Message = message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// Send Get All Rooms
func sendGetRoomsResponse(w http.ResponseWriter, status int, message string, data []m.Rooms) {
	var response m.RoomsResponse
	response.Data = data
	response.Status = status
	response.Message = message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

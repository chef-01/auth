package helpers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

// APIResponse defines the standard response structure.
type APIResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"` // Added status code to the response body.
	Message string      `json:"message"`
	Meta    interface{} `json:"meta,omitempty"` // Optional metadata.
	Data    interface{} `json:"data,omitempty"` // Actual response data.
}

// JSONResponse sends a standardized JSON response.
// It enriches the metadata with default fields recommended for RESTful APIs.
func JSONResponse(w http.ResponseWriter, status int, success bool, message string, meta interface{}, data interface{}) {
	// Get API version from environment or default to "v1".
	apiVersion := os.Getenv("API_VERSION")
	if apiVersion == "" {
		apiVersion = "v1"
	}

	// Standard metadata that will be sent with every response.
	defaultMeta := map[string]interface{}{
		"timestamp":   time.Now().Format(time.RFC3339),
		"api_version": apiVersion,
	}

	// Merge provided meta (if any) with the default meta.
	if meta != nil {
		if m, ok := meta.(map[string]interface{}); ok {
			for k, v := range m {
				defaultMeta[k] = v
			}
		} else {
			defaultMeta["detail"] = meta
		}
	}

	response := APIResponse{
		Success: success,
		Status:  status,
		Message: message,
		Meta:    defaultMeta,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

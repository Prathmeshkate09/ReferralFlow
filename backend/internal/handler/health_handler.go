package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "healthy",
		"service":   "ReferralFlow API",
		"version":   "1.0.0",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

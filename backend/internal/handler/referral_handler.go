package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/model"
	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/service"
)

type ReferralHandler struct {
	referralService *service.ReferralService
	mockProfile     *model.CandidateProfile
}

func NewReferralHandler(referralService *service.ReferralService, profile *model.CandidateProfile) *ReferralHandler {
	return &ReferralHandler{
		referralService: referralService,
		mockProfile:     profile,
	}
}

func (h *ReferralHandler) HandleCreateReferral(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input model.CreateReferralInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if input.JobID == "" {
		http.Error(w, "Missing job_id", http.StatusBadRequest)
		return
	}

	ref, err := h.referralService.RequestReferral(input, h.mockProfile)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Referral requested successfully! Matching with active referrers.",
		"data":    ref,
	})
}

func (h *ReferralHandler) HandleListReferrals(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	referrals, err := h.referralService.GetAllReferrals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"count":   len(referrals),
		"data":    referrals,
	})
}

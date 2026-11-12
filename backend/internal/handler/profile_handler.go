package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/model"
	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/service"
)

type ProfileHandler struct {
	scoringService *service.ScoringService
	profile        *model.CandidateProfile
}

func NewProfileHandler(scoringService *service.ScoringService, profile *model.CandidateProfile) *ProfileHandler {
	return &ProfileHandler{
		scoringService: scoringService,
		profile:        profile,
	}
}

func (h *ProfileHandler) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	check := h.scoringService.EvaluateProfile(h.profile)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"profile":    h.profile,
			"evaluation": check,
		},
	})
}

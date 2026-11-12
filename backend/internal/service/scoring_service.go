package service

import "github.com/Prathmeshkate09/ReferralFlow/backend/internal/model"

const ReferralUnlockThreshold = 80

// ScoringService calculates candidate profile completeness and referral eligibility.
type ScoringService struct{}

func NewScoringService() *ScoringService {
	return &ScoringService{}
}

// EvaluateProfile computes a weighted completion score (0-100) and returns unlock status.
func (s *ScoringService) EvaluateProfile(p *model.CandidateProfile) model.ProfileCompletionCheck {
	score := 0
	var missing []string

	// Basic information (20 points)
	if p.FullName != "" && p.Email != "" && p.Phone != "" {
		score += 20
	} else {
		missing = append(missing, "Full contact details (Name, Email, Phone)")
	}

	// Academic credentials (20 points)
	if p.College != "" && p.Degree != "" && p.GraduationYear > 0 {
		score += 20
	} else {
		missing = append(missing, "Academic info (College, Degree, Grad Year)")
	}

	// Verified Resume (25 points)
	if p.ResumeURL != "" {
		score += 25
	} else {
		missing = append(missing, "Uploaded PDF Resume")
	}

	// Technical presence: GitHub & LinkedIn (15 points)
	if p.GitHubURL != "" && p.LinkedInURL != "" {
		score += 15
	} else if p.GitHubURL != "" || p.LinkedInURL != "" {
		score += 8
		missing = append(missing, "Complete social presence (GitHub & LinkedIn)")
	} else {
		missing = append(missing, "Portfolio links (GitHub and LinkedIn)")
	}

	// Verified Skill Set: at least 3 skills (20 points)
	if len(p.Skills) >= 3 {
		score += 20
	} else if len(p.Skills) > 0 {
		score += 10
		missing = append(missing, "Add at least 3 verified technical skills")
	} else {
		missing = append(missing, "Add at least 3 verified technical skills")
	}

	if score > 100 {
		score = 100
	}

	p.CompletionScore = score
	p.ReferralUnlocked = score >= ReferralUnlockThreshold

	return model.ProfileCompletionCheck{
		Score:      score,
		IsUnlocked: p.ReferralUnlocked,
		Threshold:  ReferralUnlockThreshold,
		Missing:    missing,
	}
}

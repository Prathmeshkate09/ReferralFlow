package service

import (
	"testing"

	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/model"
)

func TestEvaluateProfile_FullProfile(t *testing.T) {
	service := NewScoringService()
	profile := &model.CandidateProfile{
		FullName:       "Prathmesh Kate",
		Email:          "kateprathamesh968@gmail.com",
		Phone:          "+91 9130482348",
		College:        "Dr. D. Y. Patil Institute of Technology",
		Degree:         "B.E. AI & Data Science",
		GraduationYear: 2026,
		ResumeURL:      "https://storage.presio.me/resumes/prathmesh_kate.pdf",
		GitHubURL:      "https://github.com/Prathmeshkate09",
		LinkedInURL:    "https://linkedin.com/in/prathamesh-kate-022799257",
		Skills:         []string{"Go", "React", "PostgreSQL", "Docker"},
	}

	result := service.EvaluateProfile(profile)

	if result.Score != 100 {
		t.Errorf("Expected score 100, got %d", result.Score)
	}
	if !result.IsUnlocked {
		t.Errorf("Expected referral to be unlocked for score 100")
	}
	if len(result.Missing) != 0 {
		t.Errorf("Expected 0 missing items, got %v", result.Missing)
	}
}

func TestEvaluateProfile_UnderThreshold(t *testing.T) {
	service := NewScoringService()
	profile := &model.CandidateProfile{
		FullName: "Candidate Doe",
		Email:    "candidate@example.com",
		// missing phone, resume, github
		College:        "Engineering College",
		Degree:         "B.Tech",
		GraduationYear: 2026,
		Skills:         []string{"Go"},
	}

	result := service.EvaluateProfile(profile)

	if result.Score >= ReferralUnlockThreshold {
		t.Errorf("Expected score below %d, got %d", ReferralUnlockThreshold, result.Score)
	}
	if result.IsUnlocked {
		t.Errorf("Expected referral to be locked when score is below threshold")
	}
}

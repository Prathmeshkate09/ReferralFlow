package model

import "time"

// Job represents a verified job listing in the system.
type Job struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Company      string    `json:"company"`
	Location     string    `json:"location"`
	Workplace    string    `json:"workplace"` // Remote, Hybrid, Onsite
	Experience   string    `json:"experience"`
	StipendRange string    `json:"stipend_range"`
	Department   string    `json:"department"`
	Skills       []string  `json:"skills"`
	IsVerified   bool      `json:"is_verified"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CandidateProfile represents the candidate profile and completion metrics.
type CandidateProfile struct {
	ID               string    `json:"id"`
	UserID           string    `json:"user_id"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	College          string    `json:"college"`
	Degree           string    `json:"degree"`
	GraduationYear   int       `json:"graduation_year"`
	ResumeURL        string    `json:"resume_url"`
	GitHubURL        string    `json:"github_url"`
	LinkedInURL      string    `json:"linkedin_url"`
	Skills           []string  `json:"skills"`
	CompletionScore  int       `json:"completion_score"` // 0 to 100
	ReferralUnlocked bool      `json:"referral_unlocked"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// ReferralStatus defines the state machine of a referral request.
type ReferralStatus string

const (
	StatusSubmitted          ReferralStatus = "SUBMITTED"
	StatusUnderReview        ReferralStatus = "UNDER_REVIEW"
	StatusMatchedToReferrer  ReferralStatus = "MATCHED_TO_REFERRER"
	StatusDeliveredToCompany ReferralStatus = "DELIVERED_TO_COMPANY"
	StatusRejected           ReferralStatus = "REJECTED"
	StatusExpired            ReferralStatus = "EXPIRED"
)

// ReferralRequest represents a request submitted by a qualified candidate.
type ReferralRequest struct {
	ID            string         `json:"id"`
	JobID         string         `json:"job_id"`
	JobTitle      string         `json:"job_title"`
	Company       string         `json:"company"`
	CandidateID   string         `json:"candidate_id"`
	CandidateName string         `json:"candidate_name"`
	PitchNote     string         `json:"pitch_note"`
	Status        ReferralStatus `json:"status"`
	ReferrerID    string         `json:"referrer_id,omitempty"`
	ReferrerName  string         `json:"referrer_name,omitempty"`
	SubmittedAt   time.Time      `json:"submitted_at"`
	ResolvedAt    *time.Time     `json:"resolved_at,omitempty"`
}

// CreateReferralInput represents the incoming payload to request a referral.
type CreateReferralInput struct {
	JobID       string `json:"job_id"`
	CandidateID string `json:"candidate_id"`
	PitchNote   string `json:"pitch_note"`
}

// ProfileCompletionCheck represents scoring details.
type ProfileCompletionCheck struct {
	Score      int      `json:"score"`
	IsUnlocked bool     `json:"is_unlocked"`
	Threshold  int      `json:"threshold"`
	Missing    []string `json:"missing_items"`
}

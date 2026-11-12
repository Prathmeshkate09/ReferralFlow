package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/model"
	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/repository"
)

var (
	ErrIneligibleCandidate = errors.New("candidate profile completion score is below 80%; referrals locked")
	ErrJobInvalid          = errors.New("invalid or unverified job reference")
)

type ReferralService struct {
	jobRepo      repository.JobRepository
	referralRepo repository.ReferralRepository
	scoring      *ScoringService
}

func NewReferralService(
	jobRepo repository.JobRepository,
	referralRepo repository.ReferralRepository,
	scoring *ScoringService,
) *ReferralService {
	return &ReferralService{
		jobRepo:      jobRepo,
		referralRepo: referralRepo,
		scoring:      scoring,
	}
}

// RequestReferral checks profile eligibility and dispatches the referral request.
func (s *ReferralService) RequestReferral(input model.CreateReferralInput, candidate *model.CandidateProfile) (*model.ReferralRequest, error) {
	// 1. Verify eligibility using scoring service
	check := s.scoring.EvaluateProfile(candidate)
	if !check.IsUnlocked {
		return nil, fmt.Errorf("%w (current score: %d/100)", ErrIneligibleCandidate, check.Score)
	}

	// 2. Fetch job details
	job, err := s.jobRepo.GetByID(input.JobID)
	if err != nil {
		return nil, ErrJobInvalid
	}

	// 3. Create initial referral request
	requestID := fmt.Sprintf("ref-%d", time.Now().UnixNano())
	ref := &model.ReferralRequest{
		ID:            requestID,
		JobID:         job.ID,
		JobTitle:      job.Title,
		Company:       job.Company,
		CandidateID:   candidate.ID,
		CandidateName: candidate.FullName,
		PitchNote:     input.PitchNote,
		Status:        model.StatusSubmitted,
		SubmittedAt:   time.Now(),
	}

	if err := s.referralRepo.Create(ref); err != nil {
		return nil, err
	}

	// 4. Asynchronous referral dispatch worker (simulating Presio's 24-hour matching pipeline)
	go s.dispatchReferralWorker(requestID, job.Company)

	return ref, nil
}

// dispatchReferralWorker processes referral lifecycle asynchronously via Goroutine
func (s *ReferralService) dispatchReferralWorker(requestID string, company string) {
	// Step 1: Transition to UNDER_REVIEW
	time.Sleep(1500 * time.Millisecond)
	_ = s.referralRepo.UpdateStatus(requestID, model.StatusUnderReview, "", "")
	log.Printf("[ReferralWorker] Request %s for %s moved to UNDER_REVIEW", requestID, company)

	// Step 2: Simulate matching with an active employee referrer
	time.Sleep(2000 * time.Millisecond)
	referrerID := "ref-user-884"
	referrerName := fmt.Sprintf("Senior Engineer @ %s", company)
	_ = s.referralRepo.UpdateStatus(requestID, model.StatusMatchedToReferrer, referrerID, referrerName)
	log.Printf("[ReferralWorker] Request %s successfully MATCHED to %s", requestID, referrerName)
}

func (s *ReferralService) GetCandidateReferrals(candidateID string) ([]model.ReferralRequest, error) {
	return s.referralRepo.ListByCandidate(candidateID)
}

func (s *ReferralService) GetAllReferrals() ([]model.ReferralRequest, error) {
	return s.referralRepo.ListAll()
}

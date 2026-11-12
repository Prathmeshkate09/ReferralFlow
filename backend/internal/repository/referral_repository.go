package repository

import (
	"sync"
	"time"

	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/model"
)

type ReferralRepository interface {
	Create(ref *model.ReferralRequest) error
	GetByID(id string) (*model.ReferralRequest, error)
	ListByCandidate(candidateID string) ([]model.ReferralRequest, error)
	ListAll() ([]model.ReferralRequest, error)
	UpdateStatus(id string, status model.ReferralStatus, referrerID string, referrerName string) error
}

type MemoryReferralRepository struct {
	mu        sync.RWMutex
	referrals map[string]model.ReferralRequest
}

func NewMemoryReferralRepository() *MemoryReferralRepository {
	return &MemoryReferralRepository{
		referrals: make(map[string]model.ReferralRequest),
	}
}

func (r *MemoryReferralRepository) Create(ref *model.ReferralRequest) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.referrals[ref.ID] = *ref
	return nil
}

func (r *MemoryReferralRepository) GetByID(id string) (*model.ReferralRequest, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ref, exists := r.referrals[id]
	if !exists {
		return nil, ErrReferralNotFound
	}
	return &ref, nil
}

func (r *MemoryReferralRepository) ListByCandidate(candidateID string) ([]model.ReferralRequest, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var list []model.ReferralRequest
	for _, ref := range r.referrals {
		if ref.CandidateID == candidateID {
			list = append(list, ref)
		}
	}
	return list, nil
}

func (r *MemoryReferralRepository) ListAll() ([]model.ReferralRequest, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var list []model.ReferralRequest
	for _, ref := range r.referrals {
		list = append(list, ref)
	}
	return list, nil
}

func (r *MemoryReferralRepository) UpdateStatus(id string, status model.ReferralStatus, referrerID string, referrerName string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	ref, exists := r.referrals[id]
	if !exists {
		return ErrReferralNotFound
	}

	ref.Status = status
	if referrerID != "" {
		ref.ReferrerID = referrerID
		ref.ReferrerName = referrerName
	}
	now := time.Now()
	ref.ResolvedAt = &now
	r.referrals[id] = ref
	return nil
}

package service

import (
	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/model"
	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/repository"
)

type JobService struct {
	repo repository.JobRepository
}

func NewJobService(repo repository.JobRepository) *JobService {
	return &JobService{repo: repo}
}

func (s *JobService) SearchJobs(query string, workplace string) ([]model.Job, error) {
	return s.repo.List(query, workplace)
}

func (s *JobService) GetJob(id string) (*model.Job, error) {
	return s.repo.GetByID(id)
}

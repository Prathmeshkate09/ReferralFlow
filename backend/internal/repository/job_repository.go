package repository

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/model"
)

var (
	ErrJobNotFound      = errors.New("job not found")
	ErrReferralNotFound = errors.New("referral request not found")
)

type JobRepository interface {
	List(filterQuery string, workplace string) ([]model.Job, error)
	GetByID(id string) (*model.Job, error)
	Create(job *model.Job) error
}

type MemoryJobRepository struct {
	mu   sync.RWMutex
	jobs map[string]model.Job
}

func NewMemoryJobRepository() *MemoryJobRepository {
	repo := &MemoryJobRepository{
		jobs: make(map[string]model.Job),
	}
	repo.seedInitialJobs()
	return repo
}

func (r *MemoryJobRepository) seedInitialJobs() {
	now := time.Now()
	sampleJobs := []model.Job{
		{
			ID:           "job-aura-01",
			Title:        "Software Engineering Intern",
			Company:      "Auraverse / Presio",
			Location:     "Remote, India",
			Workplace:    "Remote",
			Experience:   "0+ years",
			StipendRange: "₹15,000 / month",
			Department:   "Core Product & Engineering",
			Skills:       []string{"Go", "React", "PostgreSQL", "REST APIs", "Docker"},
			IsVerified:   true,
			CreatedAt:    now.Add(-48 * time.Hour),
			UpdatedAt:    now,
		},
		{
			ID:           "job-fintech-02",
			Title:        "Backend Systems Intern (Go)",
			Company:      "FinRoute Labs",
			Location:     "Bengaluru, India",
			Workplace:    "Hybrid",
			Experience:   "0-1 years",
			StipendRange: "₹25,000 / month",
			Department:   "Platform Engineering",
			Skills:       []string{"Go", "PostgreSQL", "Kafka", "Microservices"},
			IsVerified:   true,
			CreatedAt:    now.Add(-24 * time.Hour),
			UpdatedAt:    now,
		},
		{
			ID:           "job-cloud-03",
			Title:        "Full Stack Developer Intern",
			Company:      "NextScale AI",
			Location:     "Remote, Global",
			Workplace:    "Remote",
			Experience:   "0+ years",
			StipendRange: "₹20,000 / month",
			Department:   "Product Engineering",
			Skills:       []string{"React", "TypeScript", "Go", "PostgreSQL", "Tailwind CSS"},
			IsVerified:   true,
			CreatedAt:    now.Add(-12 * time.Hour),
			UpdatedAt:    now,
		},
		{
			ID:           "job-infra-04",
			Title:        "Junior Database & Backend Engineer",
			Company:      "DataPulse Systems",
			Location:     "Pune, India",
			Workplace:    "Onsite",
			Experience:   "0-1 years",
			StipendRange: "₹30,000 / month",
			Department:   "Infrastructure",
			Skills:       []string{"PostgreSQL", "SQL Tuning", "Go", "Redis", "Linux"},
			IsVerified:   true,
			CreatedAt:    now.Add(-6 * time.Hour),
			UpdatedAt:    now,
		},
	}

	for _, j := range sampleJobs {
		r.jobs[j.ID] = j
	}
}

func (r *MemoryJobRepository) List(filterQuery string, workplace string) ([]model.Job, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	query := strings.ToLower(strings.TrimSpace(filterQuery))
	wp := strings.ToLower(strings.TrimSpace(workplace))

	var list []model.Job
	for _, job := range r.jobs {
		if wp != "" && wp != "all" && strings.ToLower(job.Workplace) != wp {
			continue
		}
		if query != "" {
			matchesTitle := strings.Contains(strings.ToLower(job.Title), query)
			matchesCompany := strings.Contains(strings.ToLower(job.Company), query)
			matchesSkill := false
			for _, s := range job.Skills {
				if strings.Contains(strings.ToLower(s), query) {
					matchesSkill = true
					break
				}
			}
			if !matchesTitle && !matchesCompany && !matchesSkill {
				continue
			}
		}
		list = append(list, job)
	}
	return list, nil
}

func (r *MemoryJobRepository) GetByID(id string) (*model.Job, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	job, exists := r.jobs[id]
	if !exists {
		return nil, ErrJobNotFound
	}
	return &job, nil
}

func (r *MemoryJobRepository) Create(job *model.Job) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.jobs[job.ID] = *job
	return nil
}

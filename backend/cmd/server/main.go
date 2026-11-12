package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/handler"
	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/middleware"
	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/model"
	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/repository"
	"github.com/Prathmeshkate09/ReferralFlow/backend/internal/service"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 1. Initialize Repositories
	jobRepo := repository.NewMemoryJobRepository()
	referralRepo := repository.NewMemoryReferralRepository()

	// 2. Initialize Core Services
	scoringService := service.NewScoringService()
	jobService := service.NewJobService(jobRepo)
	referralService := service.NewReferralService(jobRepo, referralRepo, scoringService)

	// Mock active candidate profile (defaults to 90% completion - unlocked)
	activeProfile := &model.CandidateProfile{
		ID:               "cand-prathmesh-01",
		UserID:           "usr-kate-968",
		FullName:         "Prathmesh Kate",
		Email:            "kateprathamesh968@gmail.com",
		Phone:            "+91 9130482348",
		College:          "Dr. D. Y. Patil Institute of Technology, Pune",
		Degree:           "B.E. Artificial Intelligence & Data Science",
		GraduationYear:   2026,
		ResumeURL:        "https://github.com/Prathmeshkate09/resume.pdf",
		GitHubURL:        "https://github.com/Prathmeshkate09",
		LinkedInURL:      "https://linkedin.com/in/prathamesh-kate-022799257",
		Skills:           []string{"Go", "React", "PostgreSQL", "Docker", "REST APIs"},
		CompletionScore:  90,
		ReferralUnlocked: true,
	}

	// 3. Initialize HTTP Handlers
	jobHandler := handler.NewJobHandler(jobService)
	referralHandler := handler.NewReferralHandler(referralService, activeProfile)
	profileHandler := handler.NewProfileHandler(scoringService, activeProfile)

	// 4. Configure Routing
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", handler.HandleHealth)
	mux.HandleFunc("/api/jobs", jobHandler.HandleListJobs)
	mux.HandleFunc("/api/referrals", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			referralHandler.HandleCreateReferral(w, r)
		} else if r.Method == http.MethodGet {
			referralHandler.HandleListReferrals(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/profile", profileHandler.HandleGetProfile)

	// Wrap with Middleware
	handlerWithMiddleware := middleware.Logger(middleware.EnableCORS(mux))

	log.Printf("=====================================================")
	log.Printf(" ReferralFlow Backend Service (Go) started on port %s", port)
	log.Printf(" Endpoints:")
	log.Printf("   - GET  /api/health")
	log.Printf("   - GET  /api/jobs?q=&workplace=")
	log.Printf("   - GET  /api/referrals")
	log.Printf("   - POST /api/referrals")
	log.Printf("   - GET  /api/profile")
	log.Printf("=====================================================")

	if err := http.ListenAndServe(":"+port, handlerWithMiddleware); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

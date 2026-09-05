package main

import (
	"log"

	"github.com/XIAOR1Ck/localsewa-backend/internal/config"
	"github.com/XIAOR1Ck/localsewa-backend/internal/handlers"
	"github.com/XIAOR1Ck/localsewa-backend/internal/server"
	"github.com/XIAOR1Ck/localsewa-backend/internal/service"
	"github.com/XIAOR1Ck/localsewa-backend/pkg/firebase"
)

func main() {
    // Load configuration
    cfg := config.Load()
		firebaseClient, err := firebase.NewClient("firebase_credentials.json")
		if err != nil {
			log.Fatalf("Error creating firebase client: %s", err)
	}
		authService := service.NewAuthService(firebaseClient)
		authHandler := handlers.NewAuthHandler(authService)

    // Setup server
	srv := server.NewServer(cfg, authHandler)

    // Start server
    if err := srv.Start(); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}

package main

import (
	"context"
	"log"

	"github.com/XIAOR1Ck/localsewa-backend/internal/config"
	"github.com/XIAOR1Ck/localsewa-backend/internal/db"
	"github.com/XIAOR1Ck/localsewa-backend/internal/handlers"
	"github.com/XIAOR1Ck/localsewa-backend/internal/server"
	"github.com/XIAOR1Ck/localsewa-backend/internal/service"
	"github.com/XIAOR1Ck/localsewa-backend/pkg/firebase"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
    // Load configuration
    cfg := config.Load()
		firebaseClient, err := firebase.NewClient(cfg.FirebaseCredentials)
		if err != nil {
			log.Fatalf("Error creating firebase client: %s", err)
	}

//Database setup
	
	conn, err := pgx.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
			log.Fatal("Error initializing database")
	}
	defer conn.Close(ctx)

	//querier
	querier := db.New(conn)

	// services
	svc := service.InitServices(querier, firebaseClient)


	//Handlers	
	authHandler := handlers.NewAuthHandler(svc.AuthService)

	
    // Setup server
	srv := server.NewServer(cfg, authHandler, querier)

    // Start server
    if err := srv.Start(); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}

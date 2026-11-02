package main

import (
	"context"
	"fmt"
	"log"

	"github.com/XIAOR1Ck/localsewa-backend/internal/auth"
	"github.com/XIAOR1Ck/localsewa-backend/internal/config"
	"github.com/XIAOR1Ck/localsewa-backend/internal/db"
	"github.com/XIAOR1Ck/localsewa-backend/internal/handlers"
	"github.com/XIAOR1Ck/localsewa-backend/internal/middleware"
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
		log.Fatalf("Error creating firebase client: %v", err)
	}
	fmt.Printf("%+v", cfg)

	//Database setup

	conn, err := pgx.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Error initializing database")
	}
	defer conn.Close(ctx)

	//querier
	querier := db.New(conn)

	// JWT service
	jwtSvc := auth.NewJWTService(cfg)

	//Middleware
	middleware := middleware.NewAuthMiddleware(jwtSvc)

	// services
	svc := service.InitServices(querier, firebaseClient)

	//Handlers
	handlers := handlers.InitHandlers(svc, jwtSvc)

	// Setup server
	srv, err := server.NewServerBuilder(cfg).
		WithQuerier(querier).
		WithAuthMiddleware(middleware).
		WithHandlers(handlers).
		WithEngine().
		WithHealthCheck().
		WithBaseRoute("/api/v1").
		WithAuthRoutes().
		Build()
	if err != nil {
		log.Fatalf("failed to start server:\n%w", err)
	}

	// Start server
	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

package main

import (
    "log"

    "github.com/XIAOR1Ck/localsewa-backend/internal/config"
    "github.com/XIAOR1Ck/localsewa-backend/internal/server"
)

func main() {
    // Load configuration
    cfg := config.Load()

    // Setup server
    srv := server.NewServer(cfg)

    // Start server
    if err := srv.Start(); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}

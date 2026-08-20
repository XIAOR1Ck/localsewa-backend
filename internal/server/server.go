package server

import (
	"fmt"
	"log"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/XIAOR1Ck/localsewa-backend/internal/config"
)

type Server struct {
	engine *gin.Engine
	config *config.Config
}

func NewServer (
	cfg *config.Config,
	//add handlers here when done
) *Server {
// Set Gin mode based on environment
    if cfg.Environment == "production" {
        gin.SetMode(gin.ReleaseMode)
    }

	engine := gin.New()

	//global middleware
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	// Custom cors muddleware
	engine.Use(corsMiddleware())


// Health check endpoint
    engine.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status":    "ok",
            "timestamp": time.Now().Unix(),
        })
    })


	// Add other required things here after this like groups etc


return &Server{
        engine: engine,
        config: cfg,
    }
}

// start server
func (s *Server) Start() error {
    addr := fmt.Sprintf(":%s", s.config.ServerPort)
    log.Printf("Server starting on %s", addr)
    return s.engine.Run(addr)
}

// CORS middleware
func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
        c.Header("Access-Control-Expose-Headers", "Content-Length")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}

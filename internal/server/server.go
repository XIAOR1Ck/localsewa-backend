package server

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/XIAOR1Ck/localsewa-backend/internal/config"
	"github.com/XIAOR1Ck/localsewa-backend/internal/db"
	"github.com/XIAOR1Ck/localsewa-backend/internal/handlers"
	"github.com/XIAOR1Ck/localsewa-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine     *gin.Engine
	config     *config.Config
	handlers   *handlers.Handlers
	querier    *db.Queries
	middleware *middleware.AuthMiddleware
}

type ServerBuilder struct {
	cfg        *config.Config
	handlers   *handlers.Handlers
	querier    *db.Queries
	middleware *middleware.AuthMiddleware
	engine     *gin.Engine
	api        *gin.RouterGroup
}

// NewServerBuilder creates a builder with required dependencies.
func NewServerBuilder(cfg *config.Config) *ServerBuilder {
	return &ServerBuilder{cfg: cfg}
}

func (b *ServerBuilder) WithHandlers(h *handlers.Handlers) *ServerBuilder {
	b.handlers = h
	return b
}

func (b *ServerBuilder) WithQuerier(q *db.Queries) *ServerBuilder {
	b.querier = q
	return b
}

func (b *ServerBuilder) WithAuthMiddleware(m *middleware.AuthMiddleware) *ServerBuilder {
	b.middleware = m
	return b
}

// WithEngine sets Gin mode and creates the engine with global middleware.
func (b *ServerBuilder) WithEngine() *ServerBuilder {
	if b.cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	b.engine = gin.New()
	b.engine.Use(gin.Recovery())
	b.engine.Use(gin.Logger())
	b.engine.Use(corsMiddleware())

	return b
}

// WithHealthCheck registers the health endpoint.
func (b *ServerBuilder) WithHealthCheck() *ServerBuilder {
	b.engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "ok",
			"timestamp": time.Now().Unix(),
		})
	})
	return b
}

// WithAPIRoutes registers all versioned API routes.
func (b *ServerBuilder) WithAPIRoutes() *ServerBuilder {
	b.api = b.engine.Group("/api/v1")

	b.registerAuthRoutes()
	// b.registerUserRoutes()
	// add other route groups here

	return b
}

func (b *ServerBuilder) registerAuthRoutes() {
	auth := b.api.Group("/auth")
	{
		auth.POST("/user/register", b.handlers.AuthHandler.UserRegister)
		auth.POST("/user/login", b.handlers.AuthHandler.UserLogin)
	}
}

// Build validates the configuration and returns the Server.
func (b *ServerBuilder) Build() (*Server, error) {
	switch {
	case b.cfg == nil:
		return nil, errors.New("server builder: config is required")
	case b.handlers == nil:
		return nil, errors.New("server builder: handlers are required")
	case b.querier == nil:
		return nil, errors.New("server builder: querier is required")
	case b.engine == nil:
		return nil, errors.New("server builder: engine not initialized, call WithEngine()")
	}

	return &Server{
		engine:   b.engine,
		config:   b.cfg,
		handlers: b.handlers,
		querier:  b.querier,
	}, nil
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

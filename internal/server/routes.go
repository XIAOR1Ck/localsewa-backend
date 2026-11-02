package server

// This file is used to seperate the Routes builder from the ServerBuilder struct for cleanliness purpose
// The routes are added to the api type in the builder which is of type *gin.RouterGroup

// WithAuthRoutes Builder used to define Auth specific routes
func (b *ServerBuilder) WithAuthRoutes() *ServerBuilder {
	auth := b.api.Group("/auth")
	auth.POST("/users/login", b.handlers.AuthHandler.UserLogin)
	auth.POST("/users/register", b.handlers.AuthHandler.UserRegister)
	auth.POST("/provider/login", b.handlers.AuthHandler.ProviderLogin)
	auth.POST("/provider/register", b.handlers.AuthHandler.ProviderRegister)

	return b
}

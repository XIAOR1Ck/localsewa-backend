package handlers

import (
	"net/http"

	"github.com/XIAOR1Ck/localsewa-backend/internal/domain"
	"github.com/XIAOR1Ck/localsewa-backend/internal/handlers/middleware"
	"github.com/XIAOR1Ck/localsewa-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
 authService *service.AuthService
}

func NewAuthHandler (authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) UserRegister (c *gin.Context) {
	//Get data from request body
	var data domain.UserRegisterModel
	if err:= c.BindJSON(&data); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	user, err := h.authService.UserRegisterWEmail(c, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Generate JWT
	claims := middleware.JWTCustomClaims{
		ID: user.ID,
		Email: user.Email,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "localservice",
			Subject:  "user-auth",
		},
	}
	token, err :=  middleware.GenerateJWT(claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	res := Response{
		Success: true,
		Message: "User Created Successfully!!",
		Token: token,
		Data: user,
	} 

	c.JSON(http.StatusOK, res)

}

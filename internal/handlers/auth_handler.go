package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/XIAOR1Ck/localsewa-backend/internal/service"
	"github.com/XIAOR1Ck/localsewa-backend/internal/domain"
)

type AuthHandler struct {
 authService *service.FirebaseAuthService
}

func NewAuthHandler (authService *service.FirebaseAuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) EmailRegister (c *gin.Context) {
	//Get data from request body
	var data domain.UserRegister
	if err:= c.BindJSON(&data); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	user, err := h.authService.CreateUserWEmail(c, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)

}

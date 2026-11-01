package handlers

import (
	"net/http"

	"github.com/XIAOR1Ck/localsewa-backend/internal/auth"
	"github.com/XIAOR1Ck/localsewa-backend/internal/domain"
	"github.com/XIAOR1Ck/localsewa-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
	jwtSvc      *auth.JWTService
}

func NewAuthHandler(authService *service.AuthService, jwtSvc *auth.JWTService) *AuthHandler {
	return &AuthHandler{authService: authService, jwtSvc: jwtSvc}
}

func (h *AuthHandler) UserRegister(c *gin.Context) {
	//Get data from request body
	var data domain.UserRegisterModel
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	user, err := h.authService.UserRegisterWEmail(c, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	//Generate JWT

	token, err := h.jwtSvc.Sign(user.ID.String(), user.Email, string(user.Role))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return

	}

	res := Response{
		Success: true,
		Message: "User Created Successfully!!",
		Token:   token,
		Data:    user,
	}

	c.JSON(http.StatusOK, res)

}

func (h *AuthHandler) UserLogin(c *gin.Context) {
	type logindata struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var data logindata
	// Get values from the request
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error})
		return
	}
	user, err := h.authService.UserLogin(c, data.Email, data.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error})
	}

	//Generate JWT

	token, err := h.jwtSvc.Sign(user.ID.String(), user.Email, string(user.Role))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return

	}

	res := Response{
		Success: true,
		Message: "User login Successful!!",
		Token:   token,
		Data:    user,
	}

	c.JSON(http.StatusOK, res)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/XIAOR1Ck/localsewa-backend/internal/auth"
	"github.com/XIAOR1Ck/localsewa-backend/internal/db"
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

	var responseUser db.GetUserByEmailRow
	jsonData, _ := json.Marshal(user)
	json.Unmarshal(jsonData, &responseUser)

	res := Response{
		Success: true,
		Message: "User Created Successfully!!",
		Token:   token,
		Data:    responseUser,
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
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	user, err := h.authService.UserLogin(c, data.Email, data.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	//Generate JWT

	token, err := h.jwtSvc.Sign(user.ID.String(), user.Email, string(user.Role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
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

func (h *AuthHandler) ProviderRegister(c *gin.Context) {
	var data domain.ProviderRegisterModel
	// Get values from the request
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	provider, err := h.authService.ProviderRegisterWEmail(c, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	//Generate JWT

	token, err := h.jwtSvc.Sign(provider.ID.String(), provider.Email, "provider")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return

	}

	var responseProvider db.GetProviderByEmailRow
	jsonData, _ := json.Marshal(provider)
	json.Unmarshal(jsonData, &responseProvider)

	res := Response{
		Success: true,
		Message: "provider login Successful!!",
		Token:   token,
		Data:    responseProvider,
	}

	c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) ProviderLogin(c *gin.Context) {
	type logindata struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var data logindata
	// Get values from the request
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	provider, err := h.authService.ProviderLogin(c, data.Email, data.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	//Generate JWT

	token, err := h.jwtSvc.Sign(provider.ID.String(), provider.Email, "provider")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return

	}

	res := Response{
		Success: true,
		Message: "provider login Successful!!",
		Token:   token,
		Data:    provider,
	}

	c.JSON(http.StatusOK, res)
}

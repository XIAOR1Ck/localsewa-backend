package handlers

import (
	"net/http"
	"strconv"

	"github.com/XIAOR1Ck/localsewa-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type ProviderHandler struct {
	svc *service.ProviderService
}

func (h *ProviderHandler) GetPovidersList(c *gin.Context) {
	limit := c.DefaultQuery("limit", "100")
	offset := c.DefaultQuery("offset", "0")

	//Type conversion of the limit and offset
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
	}
	results, err := h.svc.GetProviderList(c, int32(limitInt), int32(offsetInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	response := Response{
		Success: true,
		Message: "Providers fetch successfull",
		Data:    results,
	}

	c.JSON(http.StatusOK, response)

}

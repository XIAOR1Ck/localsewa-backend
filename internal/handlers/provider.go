package handlers

import (
	"net/http"

	"github.com/XIAOR1Ck/localsewa-backend/internal/db"
	"github.com/XIAOR1Ck/localsewa-backend/internal/domain"
	"github.com/XIAOR1Ck/localsewa-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProviderHandler struct {
	svc *service.ProviderService
}

func (h *ProviderHandler) GetPovidersList(c *gin.Context) {
	var data domain.ListProvidersQuery

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
	}
	params := db.GetAllProvidersParams{
		Categories:    data.Categories,
		MinRate:       data.MinRate,
		MaxRate:       data.MaxRate,
		MinRating:     data.MinRating,
		MinExperience: data.MinExperience,
		MaxExperience: data.MaxExperience,
		Search:        data.Search,
		Limit:         data.Limit,
		Offset:        (data.Page - 1) * data.Limit,
	}

	results, err := h.svc.GetProviderList(c, params)
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

func (h *ProviderHandler) AddCategory(c *gin.Context) {
	type categoryData struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	var data categoryData
	c.BindJSON(&data)
	category := domain.NewCategory(data.Name, data.Description)
	cg := db.AddCategoryParams{
		ID:          category.ID,
		Name:        category.Name,
		Description: &category.Description,
	}
	newCategory, err := h.svc.AddCategory(c, cg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Message: "Failed to create category!!"})
		return
	}
	c.JSON(http.StatusCreated, Response{Success: true, Message: "Category created successful", Data: newCategory})

}

func (h *ProviderHandler) GetCategories(c *gin.Context) {
	categories, err := h.svc.GetCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Message: "failed to get categories"})
		return
	}
	c.JSON(http.StatusOK, Response{Success: true, Message: "Categories fetched successfully!", Data: categories})

}

func (h *ProviderHandler) UpdateProvider(c *gin.Context) {
	id := c.Param("id")
	var userData db.UpdateProviderParams
	c.BindJSON(&userData)
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Message: "invalid user id"})
	}
	userData.ID = uid
	n, err := h.svc.UpdateProvier(c, userData)
	if n == 0 {
		c.JSON(http.StatusBadRequest, Response{Success: false, Message: "Invalid User"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Success: false, Message: "failed to update user"})
		return
	}
	c.JSON(http.StatusNoContent, Response{Success: true, Message: "Provider Updated Successfully!"})
}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"github.com/hegervalesin/helplinego/internal/services"
)

type FacilityHandler struct {
	Facility *services.FacilityService
}

func NewFacilityHandler(facilitys *services.FacilityService) *FacilityHandler {
	return &FacilityHandler{Facility: facilitys}
}

type CreateFacilityRequest struct {
	Name        string         `json:"name" binding:"required"`
	Abreviation string         `json:"abreviation" binding:"required"`
	Floors      []models.Floor `json:"floors" binding:"required"`
}

func (h *FacilityHandler) Create(c *gin.Context) {
	var req CreateFacilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	facility := models.Facility{
		Name:        req.Name,
		Abreviation: req.Abreviation,
		Floors:      req.Floors,
	}
	createdFacility, err := h.Facility.Create(facility)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdFacility)
}

func (h *FacilityHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	facilityID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid facility ID"})
		return
	}
	facility, err := h.Facility.GetByID(facilityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, facility)
}

func (h *FacilityHandler) Update(c *gin.Context) {
	id := c.Param("id")
	facilityID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid facility ID"})
		return
	}
	var req CreateFacilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	facility := models.Facility{
		ID:          facilityID,
		Name:        req.Name,
		Abreviation: req.Abreviation,
		Floors:      req.Floors,
	}
	updatedFacility, err := h.Facility.Update(facility)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedFacility)
}

func (h *FacilityHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	facilityID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid facility ID"})
		return
	}
	err = h.Facility.Delete(facilityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Facility deleted successfully"})
}

func (h *FacilityHandler) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if err != nil || perPage < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid per_page number"})
		return
	}

	facilities, total, err := h.Facility.List(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": facilities,
		"meta": gin.H{
			"page":     page,
			"per_page": perPage,
			"total":    total,
		},
	})
}

func (h *FacilityHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/facility", h.Create)
	router.GET("/facility/:id", h.GetByID)
	router.PUT("/facility/:id", h.Update)
	router.DELETE("/facility/:id", h.Delete)
	router.GET("/facility/complete", h.List)
}

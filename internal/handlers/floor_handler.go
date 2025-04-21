// internal/handlers/floor_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"github.com/hegervalesin/helplinego/internal/services"
)

type FloorHandler struct {
	FloorService *services.FloorService
}

func NewFloorHandler(floorService *services.FloorService) *FloorHandler {
	return &FloorHandler{FloorService: floorService}
}

type CreateFloorRequest struct {
	Name        string        `json:"name" binding:"required"`
	FloorNumber string        `json:"floor_number" binding:"required"`
	FacilityID  uuid.UUID     `json:"facility_id" binding:"required"`
	Rooms       []models.Room `json:"rooms" binding:"required"`
}

func (h *FloorHandler) Create(c *gin.Context) {
	var req CreateFloorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	floor := models.Floor{
		Name:        req.Name,
		FloorNumber: req.FloorNumber,
		FacilityID:  req.FacilityID,
		Rooms:       []models.Room{},
	}

	createdFloor, err := h.FloorService.Create(floor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdFloor)
}

func (h *FloorHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	floorID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	floor, err := h.FloorService.GetByID(floorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Andar não encontrado"})
		return
	}

	c.JSON(http.StatusOK, floor)
}

func (h *FloorHandler) Update(c *gin.Context) {
	id := c.Param("id")
	floorID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req CreateFloorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	floor := models.Floor{
		ID:          floorID,
		Name:        req.Name,
		FloorNumber: req.FloorNumber,
		FacilityID:  req.FacilityID,
		Rooms:       []models.Room{},
	}

	updatedFloor, err := h.FloorService.Update(floor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedFloor)
}

func (h *FloorHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	floorID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.FloorService.Delete(floorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Andar excluído com sucesso"})
}

func (h *FloorHandler) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil || page < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if err != nil || perPage < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Número de itens por página inválido"})
		return
	}

	floors, total, err := h.FloorService.List(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": floors,
		"meta": gin.H{
			"page":     page,
			"per_page": perPage,
			"total":    total,
		},
	})
}

func (h *FloorHandler) RegisterRoutes(router *gin.RouterGroup) {
	floorRoutes := router.Group("/floors")
	{
		floorRoutes.POST("", h.Create)
		floorRoutes.GET("/:id", h.GetByID)
		floorRoutes.PUT("/:id", h.Update)
		floorRoutes.DELETE("/:id", h.Delete)
		floorRoutes.GET("", h.List)
	}
}

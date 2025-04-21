package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"github.com/hegervalesin/helplinego/internal/services"
)

type SectorHandler struct {
	Service *services.SectorService
}

func NewSectorHandler(service *services.SectorService) *SectorHandler {
	return &SectorHandler{Service: service}
}

type CreateSectorRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (h *SectorHandler) Create(c *gin.Context) {
	var req CreateSectorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sector := models.Sector{
		Name:        req.Name,
		Description: req.Description,
	}

	createdSector, err := h.Service.Create(sector)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdSector)
}

func (h *SectorHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	sectorID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	sector, err := h.Service.GetByID(sectorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Setor não encontrado"})
		return
	}

	c.JSON(http.StatusOK, sector)
}

func (h *SectorHandler) Update(c *gin.Context) {
	id := c.Param("id")
	sectorID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req CreateSectorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sector := models.Sector{
		ID:          sectorID,
		Name:        req.Name,
		Description: req.Description,
	}

	updatedSector, err := h.Service.Update(sector)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedSector)
}

func (h *SectorHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	sectorID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.Service.Delete(sectorID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Setor excluído com sucesso"})
}

func (h *SectorHandler) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Número de página inválido"})
		return
	}

	perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if err != nil || perPage < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Número de itens por página inválido"})
		return
	}

	sectors, total, err := h.Service.List(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": sectors,
		"meta": gin.H{
			"page":     page,
			"per_page": perPage,
			"total":    total,
		},
	})
}

func (h *SectorHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/sectors", h.Create)
	router.GET("/sectors/:id", h.GetByID)
	router.PUT("/sectors/:id", h.Update)
	router.DELETE("/sectors/:id", h.Delete)
	router.GET("/sectors", h.List)
}

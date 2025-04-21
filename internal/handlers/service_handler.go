// internal/handlers/service_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"github.com/hegervalesin/helplinego/internal/services"
)

type ServiceHandler struct {
	ServiceService *services.ServiceService
}

func NewServiceHandler(serviceService *services.ServiceService) *ServiceHandler {
	return &ServiceHandler{ServiceService: serviceService}
}

type CreateServiceRequest struct {
	Name        string               `json:"name" binding:"required"`
	Description string               `json:"description" binding:"required"`
	Status      models.ServiceStatus `json:"status" binding:"required"`
	SectorID    uuid.UUID            `json:"sector_id" binding:"required"`
}

func (h *ServiceHandler) Create(c *gin.Context) {
	var req CreateServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := models.Services{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		SectorID:    req.SectorID,
	}

	createdService, err := h.ServiceService.Create(service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdService)
}

func (h *ServiceHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	servicesID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	service, err := h.ServiceService.GetByID(servicesID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Serviço não encontrado"})
		return
	}

	c.JSON(http.StatusOK, service)
}

func (h *ServiceHandler) Update(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req CreateServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := models.Services{
		ID:          serviceID,
		Name:        req.Name,
		Description: req.Description,
		SectorID:    req.SectorID,
	}

	updatedService, err := h.ServiceService.Update(service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedService)
}

func (h *ServiceHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	servicesID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.ServiceService.Delete(servicesID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Serviço excluído com sucesso"})
}

func (h *ServiceHandler) List(c *gin.Context) {
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

	services, total, err := h.ServiceService.List(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": services,
		"meta": gin.H{
			"page":     page,
			"per_page": perPage,
			"total":    total,
		},
	})
}

func (h *ServiceHandler) RegisterRoutes(router *gin.RouterGroup) {
	serviceRoutes := router.Group("/services")
	{
		serviceRoutes.POST("", h.Create)
		serviceRoutes.GET("/:id", h.GetByID)
		serviceRoutes.PUT("/:id", h.Update)
		serviceRoutes.DELETE("/:id", h.Delete)
		serviceRoutes.GET("", h.List)
	}
}

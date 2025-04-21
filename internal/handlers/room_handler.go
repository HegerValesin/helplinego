// internal/handlers/room_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"github.com/hegervalesin/helplinego/internal/services"
)

type RoomHandler struct {
	RoomService *services.RoomService
}

func NewRoomHandler(roomService *services.RoomService) *RoomHandler {
	return &RoomHandler{RoomService: roomService}
}

type CreateRoomRequest struct {
	Name       string    `json:"name" binding:"required"`
	RoomNumber string    `json:"roomNumber" binding:"required"`
	FloorID    uuid.UUID `json:"floor_id" binding:"required"`
	SectorID   uuid.UUID `json:"sector_id" binding:"required"`
}

func (h *RoomHandler) Create(c *gin.Context) {
	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room := models.Room{
		Name:       req.Name,
		RoomNumber: req.RoomNumber,
		FloorID:    req.FloorID,
		SectorID:   req.SectorID,
	}

	createdRoom, err := h.RoomService.Create(room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdRoom)
}

func (h *RoomHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	roomID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	room, err := h.RoomService.GetByID(roomID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sala não encontrada"})
		return
	}

	c.JSON(http.StatusOK, room)
}

func (h *RoomHandler) Update(c *gin.Context) {
	id := c.Param("id")
	roomID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room := models.Room{
		ID:         roomID,
		Name:       req.Name,
		RoomNumber: req.RoomNumber,
		FloorID:    req.FloorID,
		SectorID:   req.SectorID,
	}

	updatedRoom, err := h.RoomService.Update(room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedRoom)
}

func (h *RoomHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	roomID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.RoomService.Delete(roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sala excluída com sucesso"})
}

func (h *RoomHandler) List(c *gin.Context) {
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

	rooms, total, err := h.RoomService.List(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": rooms,
		"meta": gin.H{
			"page":     page,
			"per_page": perPage,
			"total":    total,
		},
	})
}

func (h *RoomHandler) RegisterRoutes(router *gin.RouterGroup) {
	roomRoutes := router.Group("/rooms")
	{
		roomRoutes.POST("", h.Create)
		roomRoutes.GET("/:id", h.GetByID)
		roomRoutes.PUT("/:id", h.Update)
		roomRoutes.DELETE("/:id", h.Delete)
		roomRoutes.GET("", h.List)
	}
}

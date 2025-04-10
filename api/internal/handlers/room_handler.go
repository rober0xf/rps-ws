package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rps/api/internal/services"
	"rps/pkg/models"
	"strconv"
)

type RoomHandler struct {
	roomService services.RoomService
}

// kinda constructor
func NewRoomHandler(roomService services.RoomService) *RoomHandler {
	return &RoomHandler{
		roomService: roomService,
	}
}

func (h *RoomHandler) CreateRoomHandler(c *gin.Context) {
	var dto models.RoomDTO

	if err := c.ShouldBind(&dto); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// call the service to create the room
	room, err := h.roomService.CreateRoomService(c.Request.Context(), dto)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, room)
}

func (h *RoomHandler) GetRoomByAddressHandler(c *gin.Context) {
	addr := c.Param("address")

	room, err := h.roomService.GetRoomByAddressService(c.Request.Context(), addr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, room)
}

func (h *RoomHandler) GetRoomByIDHandler(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	room, err := h.roomService.GetRoomByIDService(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, room)
}

func (h *RoomHandler) ListRoomsHandler(c *gin.Context) {
	rooms, err := h.roomService.ListRoomsService(c.Request.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "no rooms found"})
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, rooms)
}

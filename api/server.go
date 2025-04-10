package api

import (
	"github.com/gin-gonic/gin"
	"rps/api/internal/handlers"
)

func InitRoutes() {
	handler := handlers.RoomHandler{}
	router := gin.Default()
	apiV1 := router.Group("/api/v1")

	apiV1.GET("/rooms", handler.ListRoomsHandler)
	apiV1.GET("/rooms/{id}", handler.GetRoomByIDHandler)
	apiV1.GET("/rooms/{address}", handler.GetRoomByAddressHandler)
	apiV1.POST("/rooms", handler.CreateRoomHandler)

	router.Run()
}

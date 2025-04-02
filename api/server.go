package api

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type RoomDTO struct {
	Name        string `json:"name" binding:"required"`
	MaxPlayers  int    `json:"max_players"`
	MinPlayers  int    `json:"min_players"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

type Room struct {
	bun.BaseModel `bun:"table:rooms,alias:r"`

	Id          int64  `bun:"id,pk,autoincrement" json:"id"`
	Name        string `bun:"name,type:varchar(128),notnull" json:"name"`
	MaxPlayers  int    `bun:"max_players,type:int,notnull" json:"max_players"`
	MinPlayers  int    `bun:"min_players,type:int,notnull" json:"min_players"`
	Description string `bun:"description,type:varchar(512),nullzero,notnull,default:''" json:"description"`

	Address string `bun:"address,type:varchar(128),notnull" json:"address"`
	Private bool   `bun:"private,type:bool,notnull,default:false" json:"private"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

func InitRoutes() {
	router := gin.Default()
	apiV1 := router.Group("/api/v1")

	apiV1.GET("/rooms", list_rooms)
	apiV1.GET("/rooms/{id}", room_by_id)
	apiV1.POST("/rooms", create_room)

	router.Run()
}

func list_rooms(ctx *gin.Context) {
	ctx.String(http.StatusOK, "should list all rooms\n")
}

func room_by_id(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "this room is id: {id}",
	})
}

func create_room(ctx *gin.Context) {
	dto := RoomDTO{}
	if err := ctx.ShouldBind(&dto); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "creating a room",
	})
}

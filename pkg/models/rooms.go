package models

import (
	"github.com/uptrace/bun"
	"time"
)

type RoomDTO struct {
	Name        string `json:"name" binding:"required"`
	MaxPlayers  int    `json:"max_players" binding:"gte1"`
	Description string `json:"description" binding:"omitempty,max=512"`
	Private     bool   `json:"private"`
}

type Room struct {
	bun.BaseModel `bun:"table:rooms,alias:r"`

	Id             uint64 `bun:"id,pk,autoincrement" json:"id"`
	Name           string `bun:"name,type:varchar(128),notnull" json:"name"`
	MaxPlayers     int    `bun:"max_players,type:int,notnull,check:max_players>0" json:"max_players"`
	CurrentPlayers int    `bun:"current_players,type:int,notnull,default:0,check:current_players>=0" json:"current_players"`
	Description    string `bun:"description,type:varchar(512),nullzero,notnull" json:"description"`

	Address string `bun:"address,type:varchar(128),notnull" json:"address"`
	Private bool   `bun:"private,type:bool,notnull,default:false" json:"private"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp,on_update:current_timestamp" json:"updated_at"`
}

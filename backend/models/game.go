package models

import (
	"cutiecat6778/dont-trust-your-friend/lib"
	"time"
)

type GameBase struct {
	ID        []byte     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (g *GameBase) Init() *lib.CustomError {
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	id, err := lib.GenerateID()
	if err != nil {
		return err
	}
	g.ID = *id

	return nil
}

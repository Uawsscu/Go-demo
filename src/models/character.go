package models

import (
	"time"

	"github.com/google/uuid"
)

type Character struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	AnimeID     uuid.UUID `json:"animeID"`
	CreateBy    string    `json:"createBy"`
	UpdateBy    string    `json:"updateBy"`
	Activate    bool      `json:"activate"`
	CreateAt    time.Time `json:"createAt"`
	UpdateAt    time.Time `json:"updateAt"`
}

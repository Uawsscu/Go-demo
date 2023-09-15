package models

import (
	"time"

	"github.com/google/uuid"
)

type Anime struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Genre    []string  `json:"genre"`
	CreateBy string    `json:"createBy"`
	UpdateBy string    `json:"updateBy"`
	Activate bool      `json:"activate"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

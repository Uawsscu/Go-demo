package models

import (
	"time"

	"github.com/google/uuid"
)

type AnimeGenre string

const (
	SHONEN  AnimeGenre = "SHONEN"
	SPORTS  AnimeGenre = "SPORTS"
	FANTASY AnimeGenre = "FANTASY"
	HORROR  AnimeGenre = "HORROR"
	OTHER   AnimeGenre = "OTHER"
)

type Anime struct {
	ID         uuid.UUID  `gorm:"default:gen_random_uuid();not null;type:uuid;primary_key;"`
	Title      string     `gorm:"not null"`
	Genre      AnimeGenre `gorm:"not null"`
	Activate   bool       `gorm:"default:true;not null"`
	CreateBy   string     `gorm:"not null"`
	UpdateBy   string     `gorm:"not null"`
	CreateAt   time.Time  `gorm:"default:now();not null;"`
	UpdateAt   time.Time  `gorm:"default:now();not null;"`
	Characters []Character
}

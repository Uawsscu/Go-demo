package models

import (
	"time"

	"github.com/google/uuid"
)

type FileInfo struct {
	ID          uuid.UUID `gorm:"default:gen_random_uuid();not null;type:uuid;primary_key;"`
	FileName    string    `gorm:"not null"`
	Extension   string    `gorm:"not null"`
	Url         string
	Description string
	InfoType    string    `gorm:"not null"`
	Activate    bool      `gorm:"default:true;not null"`
	CreateBy    string    `gorm:"not null"`
	UpdateBy    string    `gorm:"not null"`
	CreateAt    time.Time `gorm:"default:now();not null;"`
	UpdateAt    time.Time `gorm:"default:now();not null;"`
}

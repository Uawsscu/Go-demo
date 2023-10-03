package file_info

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type InfoType string

const (
	TYPE_ANIME     InfoType = "demo/animes"
	TYPE_CHARACTER InfoType = "demo/characters"
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

type FileInfoRequest struct {
	File        *multipart.FileHeader `form:"file" binding:"required"`
	FileName    string                `form:"fileName"`
	Description string                `form:"description"`
	InfoType    string                `form:"infoType" binding:"required,oneof=demo/characters demo/animes"`
	UpdateBy    string                `form:"updateBy" binding:"required"`
}

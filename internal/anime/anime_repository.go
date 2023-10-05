package anime

import (
	"errors"
	"fmt"
	"go-demo/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnimeRepository struct{}

func (ar *AnimeRepository) GetAnimeByIDGin(animeID string) (*Anime, error) {
	fmt.Println("GetAnimeByIDGin repo...")
	var anime Anime
	err := config.DB.Where("id = ?", animeID).First(&anime).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("id not found")
		}
		return nil, err
	}
	return &anime, nil
}

func (r *AnimeRepository) CheckIDExist(id uuid.UUID) error {
	var anime Anime
	if err := config.DB.Where("id = ?", id).First(&anime).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("animeId not found")
		}
		return err
	}
	return nil
}

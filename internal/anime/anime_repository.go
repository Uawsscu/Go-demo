package anime

import (
	"fmt"
	"go-demo/config"
)

type AnimeRepository struct{}

func (ar *AnimeRepository) GetAnimeByIDGin(animeID string) (*Anime, error) {
	fmt.Println("GetAnimeByIDGin repo...")
	var anime Anime
	err := config.DB.Where("id = ?", animeID).First(&anime).Error
	if err != nil {
		return nil, err
	}
	return &anime, err
}

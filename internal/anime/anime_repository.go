package anime

import (
	"fmt"
	"main/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AnimeRepository struct{}

func (ar *AnimeRepository) GetAnimeByIDGin(c *gin.Context, animeID string) (*Anime, error) {
	fmt.Println("GetAnimeByIDGin repo...")
	var anime Anime
	err := config.DB.Where("id = ?", animeID).First(&anime).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anime not found"})
	}

	return &anime, err
}

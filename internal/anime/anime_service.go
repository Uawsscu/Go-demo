package anime

import (
	"github.com/gin-gonic/gin"
)

type AnimeService struct{}

func (ar *AnimeService) GetAnimeByID(c *gin.Params) (*Anime, error) {
	animeID := c.ByName("id")

	animeRepo := &AnimeRepository{}
	anime, err := animeRepo.GetAnimeByIDGin(animeID)

	if err != nil {
		return anime, err
	}

	return anime, err

}

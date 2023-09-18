package anime

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type AnimeService struct{}

func (ar *AnimeService) GetAnimeByID(c *gin.Context) (*Anime, error) {
	// prepare request
	animeID := c.Param("id")

	animeRepo := &AnimeRepository{}
	anime, err := animeRepo.GetAnimeByIDGin(c, animeID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GetAnime Service", anime)
	return anime, nil
}

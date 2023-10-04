package anime

import (
	"log"

	"github.com/gin-gonic/gin"
)

type AnimeService struct{}

func (ar *AnimeService) GetAnimeByID(c *gin.Context) {
	animeID := c.Param("id")

	animeRepo := &AnimeRepository{}
	anime, err := animeRepo.GetAnimeByIDGin(c, animeID)

	if err != nil {
		log.Fatal(err)
	}
	c.JSON(201, anime)
}

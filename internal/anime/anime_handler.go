package anime

import (
	"fmt"
	"main/pkg/kafkautil"

	"github.com/gin-gonic/gin"
)

// GetAnimeHandler handles GET requests to retrieve an Anime by ID.
func GetAnimeHandler(c *gin.Context) {
	fmt.Println("GetAnimeHandler start....")
	animeService := &AnimeService{}
	animeService.GetAnimeByID(c)

	kafkautil.SendMessageCreateAnime("Hello")
	fmt.Println("GetAnimeHandler success....")
}

func GetAnimeByCondition(c *gin.Context) {
	fmt.Println("GetAnimeByCondition start....")

	fmt.Println("GetAnimeByCondition success....")
}

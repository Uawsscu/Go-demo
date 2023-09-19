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
	fmt.Println("GetAnimeHandler success....")
}

func TestKafkaDemoCreateHandler(c *gin.Context) {
	fmt.Println("TestKafkaHandler start....")
	kafkautil.SendMessageCreateAnime("Hello")
	fmt.Println("TestKafkaHandler success....")
}

func GetAnimeByCondition(c *gin.Context) {
	fmt.Println("GetAnimeByCondition start....")

	fmt.Println("GetAnimeByCondition success....")
}

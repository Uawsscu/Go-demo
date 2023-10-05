package anime

import (
	"fmt"
	"go-demo/pkg/kafkautil"

	"github.com/gin-gonic/gin"
)

// @Summary Get an anime by ID
// @Description Get an anime by its ID
// @Accept json
// @Produce json
// @Tags Animes
// @Param id path string true "Anime ID"
// @Success 200 {object} Anime "Successful response"
// @Router /anime/{id} [get]
func GetAnimeHandler(c *gin.Context) {
	fmt.Println("GetAnimeHandler start....")
	animeService := &AnimeService{}
	res, err := animeService.GetAnimeByID(&c.Params)
	if err != nil {
		c.Set("error", err.Error())
	}
	c.Set("response", res)
	fmt.Println("GetAnimeHandler success....")
}

// @Summary Send kafka message
// @Description Send kafka message
// @Accept json
// @Produce json
// @Tags Kafka Demo
// @Success 200 {object} Anime "Successful response"
// @Router /anime/demoKafka [get]
func TestKafkaDemoCreateHandler(c *gin.Context) {
	fmt.Println("TestKafkaHandler start....")
	kafkautil.SendMessageCreateAnime("Hello")
	fmt.Println("TestKafkaHandler success....")
}

func GetAnimeByCondition(c *gin.Context) {
	fmt.Println("GetAnimeByCondition start....")

	fmt.Println("GetAnimeByCondition success....")
}

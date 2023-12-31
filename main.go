package main

import (
	"fmt"
	"go-demo/config"
	"os"

	_ "go-demo/docs"
	"go-demo/internal/anime"
	"go-demo/internal/character"
	"go-demo/internal/elasticsearch"
	"go-demo/internal/file_info"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	config.ConnectElasticSearch()
	config.ConnectKafka()
	config.ConnectDB()
	router := gin.Default()

	config.SwaggerConfig(router)
	router.Use(config.TraceRequest)
	initRoutes(router)

	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}

func initRoutes(eg *gin.Engine) {

	eg.GET("/anime/:id", anime.GetAnimeHandler)
	eg.GET("/anime/demoKafka", anime.TestKafkaDemoCreateHandler)
	eg.POST("/characters/create", character.CreateCharacter)
	eg.POST("/elasticsearch/create-character-index", elasticsearch.CreateElasticIndexCharacter)
	eg.POST("/files/uploadFileImage", file_info.UploadFileImage)
	// r.DELETE("/anime/delete", anime.animeDELETE)
	// r.PUT("/anime/update", anime.animeUpdate)

}

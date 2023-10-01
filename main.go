package main

import (
	"fmt"
	"go-demo/config"
	"os"

	_ "go-demo/docs"
	"go-demo/internal/anime"
	"net/http"

	"github.com/gin-contrib/cors"
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

	config.ConnectKafka()
	config.ConnectDB()
	router := gin.Default()

	config.SwaggerConfig(router)

	initRoutes(router)

	router.Use(cors.Default())
	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}

func initRoutes(eg *gin.Engine) {
	eg.GET("/anime/:id", anime.GetAnimeHandler)
	eg.GET("/anime/demoKafka", anime.TestKafkaDemoCreateHandler)
	// r.POST("/anime/demoKafka", anime.TestKafkaDemoCreateHandler)
	// r.DELETE("/anime/delete", anime.animeDELETE)
	// r.PUT("/anime/update", anime.animeUpdate)
}

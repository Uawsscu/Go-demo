package main

import (
	"fmt"
	"main/config"
	"main/internal/anime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	config.ConnectKafka()
	config.ConnectDB()
	router := gin.Default()
	router.Use(cors.Default())
	initRoutes(router)
	router.Run(":8080")
}

func initRoutes(r *gin.Engine) {
	r.GET("/anime/:id", anime.GetAnimeHandler)
	r.GET("/anime/demoKafka", anime.TestKafkaDemoCreateHandler)
	// r.POST("/anime/demoKafka", anime.TestKafkaDemoCreateHandler)
	// r.DELETE("/anime/delete", anime.animeDELETE)
	// r.PUT("/anime/update", anime.animeUpdate)

}

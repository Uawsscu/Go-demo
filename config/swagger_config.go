package config

import (
	"fmt"
	"go-demo/docs"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerConfig(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Swagger Go-Demo API"
	docs.SwaggerInfo.Description = "This is a sample server demo."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/go-demo/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

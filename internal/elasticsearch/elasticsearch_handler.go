package elasticsearch

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Summary create character index for elasticsearch
// @Description create elastic character_index
// @Accept json
// @Produce json
// @Tags Elasticsearch
// @Success 200 {object} CreateElasticIndexResponse "Successful response"
// @Router /elasticsearch/create-character-index [post]
func CreateElasticIndexCharacter(c *gin.Context) {
	fmt.Println("CreateElasticIndexCharacter start....")
	elasticsearchService := &ElasticsearchService{}
	elasticsearchService.CreateElasticIndexCharacter(c)
	fmt.Println("CreateElasticIndexCharacter success....")
}

// // @Summary create character index for elasticsearch
// // @Description create elastic character_index
// // @Accept json
// // @Produce json
// // @Tags Elasticsearch
// // @Success 200 {object} CreateElasticIndexResponse "Successful response"
// // @Router /elasticsearch/getByIndex [get]
// func getAllIndexs(c *gin.Context) {
// 	fmt.Println("CreateElasticIndexCharacter start....")
// 	// elasticsearchService := &ElasticsearchService{}
// 	fmt.Println("CreateElasticIndexCharacter success....")
// }

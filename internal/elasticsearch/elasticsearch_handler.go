package elasticsearch

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status string      `json:"status"`
	Error  ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func NewErrorResponse(message string, details string) ErrorResponse {
	return ErrorResponse{
		Status: "error",
		Error: ErrorDetail{
			Message: message,
			Details: details,
		},
	}
}

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
	res, err := elasticsearchService.CreateElasticIndexCharacter()
	if err != nil {
		c.Set("error", err.Error())
	}
	c.Set("response", res)
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

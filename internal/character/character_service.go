package character

import (
	"encoding/json"
	"go-demo/internal/elasticsearch"

	"github.com/gin-gonic/gin"
)

type CharacterService struct{}

func (cs *CharacterService) CreateCharacter(c *gin.Context) {
	var inputCharacter CreateCharacterRequest
	if err := c.ShouldBindJSON(&inputCharacter); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	characterRepository := &CharacterRepository{}
	character, err := characterRepository.CreateCharacter(inputCharacter)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	elasticsearchService := &elasticsearch.ElasticsearchService{}
	characterJSON, err := json.Marshal(character)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	err = elasticsearchService.InsertDocument("character_index", character.ID.String(), string(characterJSON))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	// return
	c.JSON(201, character)
}

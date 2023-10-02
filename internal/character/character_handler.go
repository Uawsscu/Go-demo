package character

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Summary create character and insert to elasticsearch
// @Description create character and insert to elasticsearch
// @Accept json
// @Produce json
// @Tags Characters
// @Param inputCharacter body CreateCharacterRequest true "Character JSON"
// @Success 200 {object} Character "Successful response"
// @Router /characters/create [post]
func CreateCharacter(c *gin.Context) {
	fmt.Println("CreateCharacter start....")
	characterService := &CharacterService{}
	characterService.CreateCharacter(c)
	fmt.Println("CreateCharacter success....")
}

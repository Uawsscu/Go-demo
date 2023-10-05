package character

import (
	"encoding/json"
	"go-demo/internal/anime"
	"go-demo/internal/elasticsearch"
)

type CharacterService struct{}

func (cs *CharacterService) CreateCharacter(inputCharacter CreateCharacterRequest) (*Character, error) {

	characterRepository := &CharacterRepository{}
	animeRepository := &anime.AnimeRepository{}

	err := animeRepository.CheckIDExist(inputCharacter.AnimeID)
	if err != nil {
		return nil, err
	}

	character, err := characterRepository.CreateCharacter(inputCharacter)
	if err != nil {
		return nil, err
	}

	elasticsearchService := &elasticsearch.ElasticsearchService{}
	characterJSON, err := json.Marshal(character)
	if err != nil {
		return nil, err
	}
	err = elasticsearchService.InsertDocument("character_index", character.ID.String(), string(characterJSON))
	if err != nil {
		return nil, err
	}
	return character, nil
}

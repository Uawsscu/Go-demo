package character

import (
	"database/sql"
	"fmt"
	"go-demo/config"
	"time"

	"github.com/google/uuid"
)

type CharacterRepository struct {
	db *sql.DB
}

func (cr *CharacterRepository) CreateCharacter(inputCharacter CreateCharacterRequest) (*Character, error) {
	characterID := uuid.New()
	character := Character{
		ID:          characterID,
		Name:        inputCharacter.Name,
		Description: inputCharacter.Description,
		AnimeID:     inputCharacter.AnimeID,
		CreateBy:    inputCharacter.UpdateBy,
		UpdateBy:    inputCharacter.UpdateBy,
		Activate:    true,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	tx := config.DB.Create(&character)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &character, nil
}

func (cr *CharacterRepository) GetCharacterByID(characterID string) (*Character, error) {
	// Call your database query logic here to fetch Character by ID
	// Example:
	query := "SELECT id, name FROM characters WHERE id = ?"
	var character Character
	err := cr.db.QueryRow(query, characterID).Scan(&character.ID, &character.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Character not found")
		}
		return nil, err
	}

	return &character, nil
}

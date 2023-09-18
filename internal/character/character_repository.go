package character

import (
	"database/sql"
	"fmt"
)

type CharacterRepository struct {
	db *sql.DB
}

func NewCharacterRepository(db *sql.DB) *CharacterRepository {
	return &CharacterRepository{db}
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

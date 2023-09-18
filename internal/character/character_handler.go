package character

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetCharacterHandler handles GET requests to retrieve a Character by ID.
func GetCharacterHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	characterID := params["id"]

	// Call CharacterService to get Character by ID
	characterService := &CharacterService{}
	character, err := characterService.GetCharacter(characterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Marshal Character to JSON
	characterJSON, err := json.Marshal(character)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(characterJSON)
}

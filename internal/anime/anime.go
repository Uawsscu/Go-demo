package anime

import (
	"github.com/google/uuid"
)

type Anime struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Genre    string    `json:"genre"`
	Activate bool      `json:"activate"`
}

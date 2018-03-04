package models

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Star struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Location  *Location `json:"location" db:"location"`
	GameID    uuid.UUID `json:"game_id" db:"game_id"`
}

func NewStar(g *Game) *Star {
	x := rand.Float64() * g.X
	y := rand.Float64() * g.Y

	star := &Star{
		Name:     fmt.Sprintf("Star(%f,%f)", x, y),
		Location: &Location{X: x, Y: y},
		GameID:   g.ID,
	}

	return star
}

// String is not required by pop and may be deleted
func (s Star) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Stars is not required by pop and may be deleted
type Stars []Star

// String is not required by pop and may be deleted
func (s Stars) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Star) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: s.Name, Name: "Name"},
	), nil
}

package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Game struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	X         float64   `json:"x" db:"x"`
	Y         float64   `json:"y" db:"y"`
	Density   int       `json:"density" db:"density"`
}

// String is not required by pop and may be deleted
func (g Game) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// Games is not required by pop and may be deleted
type Games []Game

// String is not required by pop and may be deleted
func (g Games) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (g *Game) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: g.Name, Name: "Name"},
		&validators.IntIsPresent{Field: g.Density, Name: "Density"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (g *Game) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (g *Game) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

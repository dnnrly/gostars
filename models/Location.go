package models

import (
	"database/sql/driver"
	"fmt"

	"github.com/golang/geo/r2"
	"github.com/pkg/errors"
)

// ErrInvalidLocation indicates that the location couldn't be deserialised
var ErrInvalidLocation = errors.New("Invalid location")

// Location is a DB compatible 2d point
type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// ToPoint converts to an r2.Point
func (l *Location) ToPoint() r2.Point {
	return r2.Point{X: l.X, Y: l.Y}
}

// Value does the conversion to a string
func (l *Location) Value() (driver.Value, error) {
	data := r2.Point{X: l.X, Y: l.Y}.String()
	return data, nil
}

// Scan converts from a string or []uint8 to a Location
func (l *Location) Scan(v interface{}) error {
	var s string
	switch v.(type) {
	case string:
		s = v.(string)
	case []uint8:
		s = string(v.([]uint8))
	default:
		return ErrInvalidLocation
	}

	_, err := fmt.Sscanf(s, "(%f,%f)", &(l.X), &(l.Y))
	if err != nil {
		return err
	}

	return nil
}

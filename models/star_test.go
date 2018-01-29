package models_test

import (
	"fmt"
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/dnnrly/gostars/models"
)

func Test_Star(t *testing.T) {
	gid, _ := uuid.NewV4()
	sid, _ := uuid.NewV4()
	s := models.Star{
		ID:       sid,
		Name:     "name",
		Location: &models.Location{X: 1.0, Y: 2.0},
		GameID:   gid,
	}

	vals, err := s.Validate(nil)
	assert.Nil(t, err)
	assert.False(t, vals.HasAny())

	expected := fmt.Sprintf(
		`{
			"id":"%s",
			"game_id":"%s",
			"created_at":"0001-01-01T00:00:00Z",
			"updated_at":"0001-01-01T00:00:00Z",
			"location": {"x":1, "y":2},
			"name": "name"
		}`,
		sid,
		gid,
	)
	assert.JSONEq(t, expected, s.String())
}

func TestNewStar(t *testing.T) {
	gid, _ := uuid.NewV4()
	g := models.Game{ID: gid, X: 1.0, Y: 2.0}

	star := models.NewStar(&g)
	assert.Equal(t, gid, star.GameID)
	assert.NotEmpty(t, star.Name)
	assert.True(t, star.Location.X >= 0.0)
	assert.True(t, star.Location.X <= 1.0)
	assert.True(t, star.Location.Y >= 0.0)
	assert.True(t, star.Location.Y <= 2.0)
}

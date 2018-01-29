package actions

import (
	"fmt"

	"github.com/dnnrly/gostars/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Game)
// DB Table: Plural (games)
// Resource: Plural (Games)
// Path: Plural (/games)
// View Template Folder: Plural (/templates/games/)

// GamesResource is the resource for the Game model
type GamesResource struct {
	buffalo.Resource
}

// List gets all Games. This function is mapped to the path
// GET /games
func (v GamesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	games := &models.Games{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Games from the DB
	if err := q.All(games); err != nil {
		return errors.WithStack(err)
	}

	// Make Games available inside the html template
	c.Set("games", games)

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("games/index.html"))
}

// Show gets the data for one Game. This function is mapped to
// the path GET /games/{game_id}
func (v GamesResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Game
	game := &models.Game{}

	// To find the Game the parameter game_id is used.
	if err := tx.Find(game, c.Param("game_id")); err != nil {
		return c.Error(404, err)
	}

	// Make game available inside the html template
	c.Set("game", game)

	return c.Render(200, r.HTML("games/show.html"))
}

// New renders the form for creating a new Game.
// This function is mapped to the path GET /games/new
func (v GamesResource) New(c buffalo.Context) error {
	// Make game available inside the html template
	c.Set("game", &models.Game{})

	return c.Render(200, r.HTML("games/new.html"))
}

// Create adds a Game to the DB. This function is mapped to the
// path POST /games
func (v GamesResource) Create(c buffalo.Context) error {
	// Allocate an empty Game
	game := &models.Game{}

	// Bind game to the html form elements
	if err := c.Bind(game); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(game)
	if err != nil {
		return errors.WithStack(err)
	}

	area := game.X * game.Y
	numStars := int(area/1000.0) * game.Density
	for i := 0; i < numStars; i++ {
		s := models.NewStar(game)
		err := tx.Create(s)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	if verrs.HasAny() {
		// Make game available inside the html template
		c.Set("game", game)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("games/new.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", fmt.Sprintf("Game was created successfully with %d stars", numStars))

	// and redirect to the games index page
	return c.Redirect(302, "/games/%s", game.ID)
}

// Edit renders a edit form for a Game. This function is
// mapped to the path GET /games/{game_id}/edit
func (v GamesResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Game
	game := &models.Game{}

	if err := tx.Find(game, c.Param("game_id")); err != nil {
		return c.Error(404, err)
	}

	// Make game available inside the html template
	c.Set("game", game)
	return c.Render(200, r.HTML("games/edit.html"))
}

// Update changes a Game in the DB. This function is mapped to
// the path PUT /games/{game_id}
func (v GamesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Game
	game := &models.Game{}

	if err := tx.Find(game, c.Param("game_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Game to the html form elements
	if err := c.Bind(game); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(game)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make game available inside the html template
		c.Set("game", game)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("games/edit.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Game was updated successfully")

	// and redirect to the games index page
	return c.Redirect(302, "/games/%s", game.ID)
}

// Destroy deletes a Game from the DB. This function is mapped
// to the path DELETE /games/{game_id}
func (v GamesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Game
	game := &models.Game{}

	// To find the Game the parameter game_id is used.
	if err := tx.Find(game, c.Param("game_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(game); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Game was destroyed successfully")

	// Redirect to the games index page
	return c.Redirect(302, "/games")
}
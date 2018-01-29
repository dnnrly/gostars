package main

import (
	"log"

	"github.com/dnnrly/gostars/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}

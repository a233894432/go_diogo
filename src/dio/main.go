package main

import (
	"github.com/ivpusic/neo"
)

func main() {
	app := neo.App()

	app.Get("/", func(ctx *neo.Ctx) (int, error) {
		return 200, ctx.Res.Text("I am Neo Programmer")
	})

	app.Start()
}

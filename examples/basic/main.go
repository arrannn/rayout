package main

import (
	"fmt"

	"github.com/arrannn/rayout"
	"github.com/arrannn/rayout/components"
)

func main() {
	app := rayout.NewApp(800, 450, "rayout User Interface")
	defer app.Close()

	var SomeState string = "Type here..."

	app.Run(
		components.Column(
			components.Label("This is a label"),
			components.VerticalSpacer(5),
			components.Button("Click me", func() { fmt.Println("Clicked") }),
			components.VerticalSpacer(5),
			components.InputText(&SomeState, 200),
		),
	)
}

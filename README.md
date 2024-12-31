# rayout
A lightweight cross-platform UI toolkit for Go powered by Raylib

```go
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
			components.TextButton("Click me", func() { fmt.Println("Clicked") }),
			components.VerticalSpacer(5),
			components.NewTextInput(&SomeState, 200),
		),
	)
}
```

## License

MIT License - see the [LICENSE](LICENSE) file for details.

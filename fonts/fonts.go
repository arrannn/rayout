package fonts

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed roboto/*.ttf
var fontFS embed.FS

func Load(f string) rl.Font {
	fontData, err := fontFS.ReadFile(f)
	if err != nil {
		panic("Failed to read embedded font: " + err.Error())
	}
	return rl.LoadFontFromMemory(".ttf", fontData, 32, nil)
}

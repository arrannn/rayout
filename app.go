package rayout

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type App struct {
	*Layout
}

func NewApp(width, height int32, title string) *App {
	rl.InitWindow(800, 450, "Ui")
	rl.SetTargetFPS(60)
	layout := NewLayout(rl.NewRectangle(0, 0, float32(width), float32(height)))
	layout.Content = &Box{
		PaddingTop:      10,
		PaddingBottom:   10,
		PaddingRight:    10,
		PaddingLeft:     10,
		BackgroundColor: &GetGlobalTheme().Background,
	}
	return &App{layout}
}

func (a *App) Run(content Widget) {
	a.Content.(*Box).Child = content
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		a.UpdateAndDraw()
		rl.EndDrawing()
	}
}

func (a *App) Close() {
	rl.CloseWindow()
}

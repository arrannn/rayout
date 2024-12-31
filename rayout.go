package rayout

import (
	"github.com/arrannn/rayout/themes"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var globalTheme *themes.Theme

func SetGlobalTheme(theme *themes.Theme) {
	globalTheme = theme
}

func GetGlobalTheme() *themes.Theme {
	return globalTheme
}

type Widget interface {
	Update(bounds rl.Rectangle) rl.Rectangle
	Draw()
}

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

type Layout struct {
	rect    rl.Rectangle
	Content Widget
}

func NewLayout(rect rl.Rectangle) *Layout {
	if globalTheme == nil {
		SetGlobalTheme(themes.RayoutTheme())
	}
	return &Layout{rect: rect}
}

func (l *Layout) Update() {
	l.Content.Update(l.rect)
}

func (l *Layout) Draw() {
	l.Content.Draw()
}

func (l *Layout) UpdateAndDraw() {
	l.Update()
	l.Draw()
}

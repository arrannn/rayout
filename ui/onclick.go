package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type OnClick struct {
	Action func()
	Child  Widget
}

func (o *OnClick) Update(bounds rl.Rectangle) rl.Rectangle {
	bounds = o.Child.Update(bounds)
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), bounds) {
			o.Action()
		}
	}
	return o.Child.Update(bounds)
}

func (o *OnClick) Draw() {
	o.Child.Draw()
}

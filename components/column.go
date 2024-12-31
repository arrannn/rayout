package components

import "github.com/arrannn/rayout/ui"

func Column(children ...ui.Widget) *ui.Column {
	return &ui.Column{Children: children}
}

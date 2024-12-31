package components

import "github.com/arrannn/rayout"

func Column(children ...rayout.Widget) *rayout.Column {
	return &rayout.Column{Children: children}
}

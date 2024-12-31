package components

import "github.com/arrannn/rayout/ui"

func Spacer(width, height float32) *ui.Spacer {
	return &ui.Spacer{Width: width, Height: height}
}

func HorizontalSpacer(width float32) *ui.Spacer {
	return &ui.Spacer{Width: width}
}

func VerticalSpacer(height float32) *ui.Spacer {
	return &ui.Spacer{Height: height}
}

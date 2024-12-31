package components

import "github.com/arrannn/rayout"

func Spacer(width, height float32) *rayout.Spacer {
	return &rayout.Spacer{Width: width, Height: height}
}

func HorizontalSpacer(width float32) *rayout.Spacer {
	return &rayout.Spacer{Width: width}
}

func VerticalSpacer(height float32) *rayout.Spacer {
	return &rayout.Spacer{Height: height}
}

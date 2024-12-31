package themes

import (
	"image/color"

	"github.com/arrannn/rayout/fonts"
)

func RayoutTheme() *Theme {
	return &Theme{
		// Fonts
		FontRegular: Font{
			Font:    fonts.Load("roboto/Roboto-Regular.ttf"),
			Size:    20,
			Spacing: 1,
		},
		FontBold: Font{
			Font:    fonts.Load("roboto/Roboto-Bold.ttf"),
			Size:    20,
			Spacing: 1,
		},

		// Text colors
		TextPrimary:   color.RGBA{R: 229, G: 229, B: 229, A: 255}, // Light gray
		TextSecondary: color.RGBA{R: 176, G: 176, B: 176, A: 255}, // Medium gray
		TextDisabled:  color.RGBA{R: 128, G: 128, B: 128, A: 255}, // Darker gray
		TextInverse:   color.RGBA{R: 33, G: 33, B: 33, A: 255},    // Near black

		// UI colors
		Primary:   color.RGBA{R: 130, G: 170, B: 255, A: 255}, // Soft blue
		Secondary: color.RGBA{R: 187, G: 134, B: 252, A: 255}, // Soft purple
		Success:   color.RGBA{R: 81, G: 207, B: 102, A: 255},  // Soft green
		Warning:   color.RGBA{R: 255, G: 179, B: 64, A: 255},  // Soft orange
		Error:     color.RGBA{R: 255, G: 85, B: 85, A: 255},   // Soft red

		// Background colors
		Background:    color.RGBA{R: 25, G: 25, B: 25, A: 255}, // Very dark gray
		BackgroundAlt: color.RGBA{R: 37, G: 37, B: 37, A: 255}, // Slightly lighter
		Surface:       color.RGBA{R: 70, G: 70, B: 70, A: 255}, // Surface elements

		// Border colors
		BorderLight:  color.RGBA{R: 100, G: 100, B: 100, A: 255}, // Subtle borders
		BorderMedium: color.RGBA{R: 73, G: 73, B: 73, A: 255},    // Standard borders
		BorderDark:   color.RGBA{R: 0, G: 0, B: 0, A: 255},       // Dark Border
		BorderFocus:  color.RGBA{R: 130, G: 170, B: 255, A: 255}, // Same as primary

		// Interactive states
		Hover:    color.RGBA{R: 255, G: 255, B: 255, A: 15}, // 6% white
		Active:   color.RGBA{R: 255, G: 255, B: 255, A: 25}, // 10% white
		Disabled: color.RGBA{R: 255, G: 255, B: 255, A: 10}, // 4% white
	}
}

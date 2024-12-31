package themes

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Font struct {
	Font    rl.Font
	Size    float32
	Spacing float32
}

type Theme struct {
	//Font
	FontRegular Font
	FontBold    Font

	// Text colors
	TextPrimary   color.RGBA // Main text color (e.g., black or dark gray)
	TextSecondary color.RGBA // Less important text (e.g., gray)
	TextDisabled  color.RGBA // Disabled text (e.g., light gray)
	TextInverse   color.RGBA // Text on dark backgrounds (e.g., white)

	// UI colors
	Primary   color.RGBA // Main accent (e.g., blue)
	Secondary color.RGBA // Secondary accent (complementary to primary)
	Success   color.RGBA // Positive actions/status (green)
	Warning   color.RGBA // Cautionary status (yellow/orange)
	Error     color.RGBA // Error states/destructive actions (red)

	// Background colors
	Background    color.RGBA // Main background (usually white or very light gray)
	BackgroundAlt color.RGBA // Alternative background (e.g., for contrast)
	Surface       color.RGBA // Card/component backgrounds

	// Border colors
	BorderLight  color.RGBA // Subtle borders (light gray)
	BorderMedium color.RGBA // Standard borders (medium gray)
	BorderDark   color.RGBA // Dark Borders (black)
	BorderFocus  color.RGBA // Focus state borders (often primary color)

	// Interactive states
	Hover    color.RGBA // Hover state color
	Active   color.RGBA // Pressed/active state
	Disabled color.RGBA // Disabled component background
}

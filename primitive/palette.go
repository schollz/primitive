package primitive

import (
	"fmt"
	"math"
)

// ColorPalette allows you to restrict
// to a certain range of colors
type ColorPalette struct {
	nameStrings []string
	hexStrings  []string
	rgbColors   []Color
}

// NewColorPalette returns a new color palette
func NewColorPalette(hexes []string, names ...string) (cp *ColorPalette, err error) {
	cp = new(ColorPalette)
	cp.hexStrings = hexes
	if len(names) > 0 {
		cp.nameStrings = names
		if len(names) != len(hexes) {
			err = fmt.Errorf("number of hex strings (%d) does not equal number of names (%d)", len(hexes), len(names))
		}
	} else {
		cp.nameStrings = hexes
	}

	cp.rgbColors = make([]Color, len(hexes))
	for i, hex := range hexes {
		cp.rgbColors[i] = MakeHexColor(hex)
	}
	return
}

// ClosestColor returns the closest RGB color in the current palette
func (cp *ColorPalette) ClosestColor(c Color) (closestColor Color, closestHex string, closestName string, err error) {
	closestI := 0
	bestScore := 100000.0
	for i, rgb := range cp.rgbColors {
		score := math.Pow(float64(rgb.R-c.R), 2)
		if score < bestScore {
			bestScore = score
			closestI = i
		}
	}
	closestColor = cp.rgbColors[closestI]
	closestHex = cp.hexStrings[closestI]
	closestName = cp.nameStrings[closestI]
	return
}

package primitive

import (
	"fmt"

	colorful "github.com/lucasb-eyer/go-colorful"
)

// ColorPalette allows you to restrict
// to a certain range of colors
type ColorPalette struct {
	nameStrings []string
	hexStrings  []string
	rgbColors   []Color
}

// NewColorPalette returns a new color palette
func NewColorPalette(hexes []string, names []string) (cp *ColorPalette, err error) {
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
	cMain := colorful.Color{float64(c.R) / 255.0, float64(c.G) / 255.0, float64(c.B) / 255.0}
	for i, c2 := range cp.rgbColors {
		cPalette := colorful.Color{float64(c2.R) / 255.0, float64(c2.G) / 255.0, float64(c2.B) / 255.0}
		score := cMain.DistanceCIEDE2000(cPalette)
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

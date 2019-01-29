package primitive

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var greenColors = strings.Split(`lawngreen 	#7CFC00 	rgb(124,252,0)
chartreuse 	#7FFF00 	rgb(127,255,0)
limegreen 	#32CD32 	rgb(50,205,50)
lime 	#00FF00 	rgb(0.255.0)
forestgreen 	#228B22 	rgb(34,139,34)
green 	#008000 	rgb(0,128,0)
darkgreen 	#006400 	rgb(0,100,0)
greenyellow 	#ADFF2F 	rgb(173,255,47)
yellowgreen 	#9ACD32 	rgb(154,205,50)
springgreen 	#00FF7F 	rgb(0,255,127)
mediumspringgreen 	#00FA9A 	rgb(0,250,154)
lightgreen 	#90EE90 	rgb(144,238,144)
palegreen 	#98FB98 	rgb(152,251,152)
darkseagreen 	#8FBC8F 	rgb(143,188,143)
mediumseagreen 	#3CB371 	rgb(60,179,113)
lightseagreen 	#20B2AA 	rgb(32,178,170)
seagreen 	#2E8B57 	rgb(46,139,87)
olive 	#808000 	rgb(128,128,0)
darkolivegreen 	#556B2F 	rgb(85,107,47)
olivedrab 	#6B8E23 	rgb(107,142,35)`, "\n")

func TestColorPalette(t *testing.T) {
	greens := make([]string, len(greenColors))
	greensName := make([]string, len(greenColors))
	for i, line := range greenColors {
		fields := strings.Fields(line)
		greens[i] = fields[1]
		greensName[i] = fields[0]
	}
	cp, err := NewColorPalette(greens, greensName...)
	assert.Nil(t, err)
	_, _, name, err := cp.ClosestColor(Color{107, 132, 35, 255})
	assert.Nil(t, err)
	assert.Equal(t, "olivedrab", name)
	_, _, name, err = cp.ClosestColor(Color{105, 130, 36, 255})
	assert.Nil(t, err)
	assert.Equal(t, "olivedrab", name)
}

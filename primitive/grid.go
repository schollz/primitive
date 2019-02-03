package primitive

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strings"
)

// DrawGrid will draw the specified number of rows and cols onto the file
func DrawGrid(rows, cols int, fname, outfilename string) (err error) {
	infile, err := os.Open(fname)
	if err != nil {
		return
	}

	// determine what kind of image it is
	var src image.Image
	if strings.HasSuffix(fname, ".jpg") {
		src, err = jpeg.Decode(infile)
	} else if strings.HasSuffix(fname, ".png") {
		src, err = png.Decode(infile)
	} else {
		src, _, err = image.Decode(infile)
	}
	infile.Close()
	if err != nil {
		return
	}

	// copy over the image, leave pixels white where you want the lines
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	fmt.Println(w, h)
	gridded := image.NewRGBA(image.Rectangle{Max: image.Point{X: w, Y: h}})
	ymod := math.Round(float64(h) / float64(rows))
	xmod := math.Round(float64(w) / float64(cols))
	fmt.Println(xmod, ymod)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := src.At(x, y)
			if math.Mod(float64(y), ymod) == 0 {
				c = color.White
			} else if math.Mod(float64(x), xmod) == 0 {
				c = color.White
			}
			gridded.Set(x, y, c)
		}
	}

	// Encode the grayscale image to the output file
	ext := strings.ToLower(filepath.Ext(outfilename))
	outfile, err := os.Create(outfilename)
	if err != nil {
		return
	}
	defer outfile.Close()
	switch ext {
	default:
		err = fmt.Errorf("unrecognized file extension: %s", ext)
	case ".png":
		err = png.Encode(outfile, gridded)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(outfile, gridded, &jpeg.Options{Quality: 90})
	}
	return
}

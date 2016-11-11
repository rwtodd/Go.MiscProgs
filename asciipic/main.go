// A utility to convert an image to ascii, at a given width (72 by default)
package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"

	"github.com/nfnt/resize"
)

// loads the image from fname, and resizes it proportionally to the given width
func loadImg(fname string, width uint) (image.Image, error) {
	rdr, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer rdr.Close()

	orig, _, err := image.Decode(rdr)
	if err != nil {
		return nil, err
	}

	ratio := float64(orig.Bounds().Dx()) / float64(orig.Bounds().Dy())
	return resize.Resize(width, uint(float64(width)/ratio), orig, resize.Bicubic), nil
}

// determine the brightness of a color, in the range 0 .. 65536
func brightness(c color.Color) float64 {
	r, g, b, _ := c.RGBA()
	return float64(r)*0.2126 + float64(g)*0.7152 + float64(b)*0.0722
}

// select a character based on a given brightness
var allchars = "#A@%$+=*:,. "

func selectChar(b float64) byte {
	return allchars[int(b*float64(len(allchars))/65536.0)]
}

// converts an image, pixel-by-pixel, to ascii
func convertImage(im image.Image) string {
	wid, ht := im.Bounds().Dx(), im.Bounds().Dy()
	var buf bytes.Buffer

	for y := 0; y < ht; y++ {
		for x := 0; x < wid; x++ {
			buf.WriteByte(selectChar(brightness(im.At(x, y))))
		}
		buf.WriteRune('\n')
	}

	return buf.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: asciipic fname width")
		return
	}

	fname := os.Args[1]
	var wid uint = 72
	if len(os.Args) > 2 {
		widat, _ := strconv.Atoi(os.Args[2])
		if widat > 0 {
			wid = uint(widat)
		}
	}

	img, err := loadImg(fname, wid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	fmt.Print(convertImage(img))
}

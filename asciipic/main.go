// A utility to convert an image to ascii, at a given width (72 by default)
package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
)

// loads the image from fname, and resizes it proportionally to the given width
func loadImg(fname string, width uint, ar float64) (image.Image, error) {
	rdr, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer rdr.Close()

	orig, _, err := image.Decode(rdr)
	if err != nil {
		return nil, err
	}

	height :=  uint((float64(width) / ar / float64(orig.Bounds().Dx())) * float64(orig.Bounds().Dy()))
	return resize.Resize(width, height, orig, resize.Bicubic), nil
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
	var flgReversed = flag.Bool("wob", false, "reverse video for white-on-black terminals")
	var flgWidth    = flag.Uint("w", 72, "desired width of output")
	var flgAspect   = flag.Float64("ar", 2.0, "aspect ratio of text (w/h)")
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		return
	}
	if *flgReversed {
		var rev = make([]byte,len(allchars))
		for i := 0 ; i < len(rev) ; i++ {
			rev[i] = allchars[len(rev)-i-1]
		}
		allchars = string(rev)
 	}
	img, err := loadImg(flag.Arg(0), *flgWidth, *flgAspect)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	fmt.Print(convertImage(img))
}

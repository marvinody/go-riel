package main

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func paintRect(rect image.Rectangle, colorIdx uint8, img *image.Paletted) {
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			img.SetColorIndex(x, y, colorIdx)
		}
	}
}

func main() {

	// Read the font data.
	fontBytes, err := ioutil.ReadFile("DTM-Mono.ttf")
	if err != nil {
		log.Println(err)
		return
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	bounds := image.Rect(0, 0, 578, 152)
	palette := []color.Color{
		color.Black,
		color.White,
	}

	img := image.NewPaletted(bounds, palette)

	borderThickness := 6

	topRow := image.Rect(0, 0, bounds.Max.X, borderThickness)
	bottomRow := image.Rect(0, bounds.Max.Y-borderThickness, bounds.Max.X, bounds.Max.Y)
	leftColumn := image.Rect(0, borderThickness, borderThickness, bounds.Max.Y-borderThickness)
	rightColumn := image.Rect(bounds.Max.X-borderThickness, borderThickness, bounds.Max.X, bounds.Max.Y-borderThickness)

	paintRect(topRow, 1, img)
	paintRect(bottomRow, 1, img)

	paintRect(leftColumn, 1, img)
	paintRect(rightColumn, 1, img)

	d := &font.Drawer{
		Dst: img,
		Src: image.White,
		Face: truetype.NewFace(f, &truetype.Options{
			Size: 24,
		}),
	}

	d.Dot = fixed.P(
		borderThickness*5,
		borderThickness*8,
	)

	d.DrawString("* hello htf, this was written in go")
	file, _ := os.Create("test.png")

	png.Encode(file, img)

}

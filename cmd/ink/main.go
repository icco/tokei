package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"github.com/fogleman/gg"
	"github.com/icco/tokei/lib"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/experimental/devices/inky"
)

var (
	Red   = color.RGBA{R: 255}
	White = color.White
	Black = color.Black
)

func main() {
	if err := lib.HostInit(); err != nil {
		log.Fatal(err)
	}

	b, err := spireg.Open("SPI0.0")
	if err != nil {
		log.Fatal(err)
	}

	dc := gpioreg.ByName("22")
	reset := gpioreg.ByName("27")
	busy := gpioreg.ByName("17")

	dev, err := inky.New(b, dc, reset, busy, &inky.Opts{
		Model:       inky.WHAT,
		ModelColor:  inky.Red,
		BorderColor: inky.White,
	})
	if err != nil {
		log.Fatalf("inky new: %+v", err)
	}

	img, err := generateImage(dev.Bounds())
	if err != nil {
		log.Fatalf("generate image: %+v", err)
	}

	if err := dev.Draw(img.Bounds(), img, image.ZP); err != nil {
		log.Fatalf("draw: %+v", err)
	}
}

func generateImage(r image.Rectangle) (image.Image, error) {
	dc := gg.NewContextForRGBA(image.NewRGBA(r))
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	fontSize := 24
	if err := dc.LoadFontFace("Roboto-Black.ttf", float64(fontSize)); err != nil {
		return nil, err
	}

	dc.SetRGB(0, 0, 0)
	lines, err := lib.GetNews(os.Getenv("NEWSAPI_KEY"), 1)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		x := 10.0
		y := 10.0
		ax := 0.5
		ay := 0.5
		width := float64(dc.Width()) - 10.0
		lineSpacing := 1.0

		dc.DrawStringWrapped(line, x, y, ax, ay, width, lineSpacing, gg.AlignLeft)
	}

	return dc.Image(), nil
}

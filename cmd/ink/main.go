package main

import (
	"image"
	"image/color"
	"log"
	"math"

	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/experimental/devices/inky"
	"periph.io/x/periph/host"
)

func main() {
	const width = 130
	const height = 50

	img := image.NewGray(image.Rectangle{Max: image.Point{X: width, Y: height}})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			dist := math.Sqrt(math.Pow(float64(x-width/2), 2)/3+math.Pow(float64(y-height/2), 2)) / (height / 1.5) * 255
			var gray uint8
			if dist > 255 {
				gray = 255
			} else {
				gray = uint8(dist)
			}
			img.SetGray(x, y, color.Gray{Y: 255 - gray})
		}
	}

	state, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Prints the loaded driver.
	log.Printf("Using drivers:\n")
	for _, driver := range state.Loaded {
		log.Printf("- %s\n", driver)
	}

	// Prints the driver that were skipped as irrelevant on the platform.
	log.Printf("Drivers skipped:\n")
	for _, failure := range state.Skipped {
		log.Printf("- %s: %s\n", failure.D, failure.Err)
	}

	// Having drivers failing to load may not require process termination. It
	// is possible to continue to run in partial failure mode.
	log.Printf("Drivers failed to load:\n")
	for _, failure := range state.Failed {
		log.Printf("- %s: %v\n", failure.D, failure.Err)
	}

	b, err := spireg.Open("SPI0.0")
	if err != nil {
		log.Fatal(err)
	}

	dc := gpioreg.ByName("22")
	reset := gpioreg.ByName("27")
	busy := gpioreg.ByName("17")

	dev, err := inky.New(b, dc, reset, busy, &inky.Opts{
		Model:       inky.PHAT,
		ModelColor:  inky.Red,
		BorderColor: inky.Black,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := dev.Draw(img.Bounds(), img, image.ZP); err != nil {
		log.Fatal(err)
	}
}

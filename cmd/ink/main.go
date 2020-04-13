package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"

	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/experimental/devices/inky"
	"periph.io/x/periph/host"
)

func main() {
	// Open and decode the image.
	f, err := os.Open("./bzl.jpg")
	if err != nil {
		log.Fatalf("couldn't open file: %+v", err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatalf("couldn't decode file: %+v", err)
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

	// Enumerate all SPI ports available and the corresponding pins.
	log.Print("SPI ports available:\n")
	for _, ref := range spireg.All() {
		log.Printf("- %s\n", ref.Name)
		if ref.Number != -1 {
			log.Printf("  %d\n", ref.Number)
		}
		if len(ref.Aliases) != 0 {
			log.Printf("  %s\n", strings.Join(ref.Aliases, " "))
		}

		p, err := ref.Open()
		if err != nil {
			log.Printf("  Failed to open: %v", err)
		}
		if p, ok := p.(spi.Pins); ok {
			log.Printf("  CLK : %s", p.CLK())
			log.Printf("  MOSI: %s", p.MOSI())
			log.Printf("  MISO: %s", p.MISO())
			log.Printf("  CS  : %s", p.CS())
		}
		if err := p.Close(); err != nil {
			log.Printf("  Failed to close: %v", err)
		}
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
		log.Fatalf("inky new: %+v", err)
	}

	if err := dev.Draw(img.Bounds(), img, image.ZP); err != nil {
		log.Fatalf("draw: %+v", err)
	}
}

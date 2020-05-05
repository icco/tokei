package lib

import (
	"fmt"
	"log"
	"strings"

	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

func HostInit() error {
	state, err := host.Init()
	if err != nil {
		return fmt.Errorf("host init: %w", err)
	}

	// Prints the loaded driver.
	log.Printf("Using drivers:\n")
	for _, driver := range state.Loaded {
		log.Printf("- %s\n", driver)
	}

	// Prints the driver that were skipped as irrelevant on the
	// platform.
	log.Printf("Drivers skipped:\n")
	for _, failure := range state.Skipped {
		log.Printf("- %s: %s\n", failure.D, failure.Err)
	}

	// Having drivers failing to load may not require process
	// termination. It is possible to continue to run in partial
	// failure mode.
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

	// Enumerate all I²C buses available and the corresponding pins.
	log.Print("I²C buses available:\n")
	for _, ref := range i2creg.All() {
		log.Printf("- %s", ref.Name)
		if ref.Number != -1 {
			log.Printf("  %d", ref.Number)
		}
		if len(ref.Aliases) != 0 {
			log.Printf("  %s", strings.Join(ref.Aliases, " "))
		}

		b, err := ref.Open()
		if err != nil {
			log.Printf("  Failed to open: %v", err)
		}
		if p, ok := b.(i2c.Pins); ok {
			log.Printf("  SDA: %s", p.SDA())
			log.Printf("  SCL: %s", p.SCL())
		}
		if err := b.Close(); err != nil {
			log.Printf("  Failed to close: %v", err)
		}
	}

	return nil
}

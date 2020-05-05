package main

import (
	"log"

	"github.com/icco/tokei/lib"
	scrollphathd "github.com/tomnz/scroll-phat-hd-go"
	"periph.io/x/periph/conn/i2c/i2creg"
)

func main() {
	if err := lib.HostInit(); err != nil {
		log.Fatal(err)
	}

	bus, err := i2creg.Open("1")
	if err != nil {
		log.Fatal(err)
	}

	display, err := scrollphathd.New(bus)
	if err != nil {
		log.Fatal(err)
	}

	display.SetBrightness(127)
	display.Fill(0, 0, 5, 5, 255)
	display.Show()
}

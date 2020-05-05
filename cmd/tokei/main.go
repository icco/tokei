package main

import (
	"log"
	"time"

	scrollphathd "github.com/icco/scroll-phat-hd-go"
	"github.com/icco/tokei/lib"
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

	for {
		display.Clear()
		display.SetBrightness(127)
		display.SetPixel(0, 0, 255)
		display.Show()
		time.Sleep(time.Second)
	}
}

package main

import (
	"log"
	"time"

	scrollphathd "github.com/icco/scroll-phat-hd-go"
	"github.com/icco/tokei/lib"
	"periph.io/x/periph/conn/i2c/i2creg"
)

var (
	height = 7
	width  = 17
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

	h := height - 1
	w := 0
	for {
		display.Clear()
		display.SetPixel(w, h, 128)
		display.Show()
		time.Sleep(time.Second)
		sec := (time.Now().Second() % 60) / 59.0
		w = w * sec
	}
}

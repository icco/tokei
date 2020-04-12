package main

import (
  	"log"
	"os"

  	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/experimental/devices/inky"
	"periph.io/x/periph/host"
)

func main() {
  if _, err := host.Init(); err != nil {
	log.Fatal(err)
}

b, err := spi.Open("SPI0.0")
if err != nil {
	log.Fatal(err)
}

dc := gpio.ByName("22")
reset := gpio.ByName("27")
busy := gpio.ByName("17")

dev, err := inky.New(b, dc, reset, busy, &inky.Opts{
	Model:       inky.PHAT,
	ModelColor:  inky.Red,
	BorderColor: inky.Black,
})
if err != nil {
	log.Fatal(err)
}

}

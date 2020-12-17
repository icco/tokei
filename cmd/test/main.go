package main

import (
	"fmt"
	"log"
	"os"

	"github.com/icco/tokei/lib"
)

func main() {
	lines, err := lib.GetNews(os.Getenv("NEWSAPI_KEY"), 3)
	if err != nil {
		log.Printf("news error: %+v", err)
	}

	for _, line := range lines {
		fmt.Printf("%s\n", line)
	}

}

package lib

import (
	"strings"
	"time"
)

func DateLines() []string {
	lines := strings.Split(time.Now().Format("15:04 Monday 2006-01-02"), " ")

	return lines
}

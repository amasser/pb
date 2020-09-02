package app

import (
	"fmt"
	"math/rand"
)

const reset = "\u001b[0m"

var colors [8]string = [8]string{
	"\x1b[92m",
	"\x1b[97m",
	"\x1b[90m",
	"\x1b[93m",
	"\x1b[94m",
	"\x1b[95m",
	"\x1b[96m",
	"\x1b[91m",
}

// TextRainbow - Do Text Like Rainbow
func TextRainbow(text string) string {
	return fmt.Sprintf("%s%s%s", colors[rand.Intn(8)], text, reset)
}

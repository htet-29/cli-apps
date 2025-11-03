package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

func main() {

	info, _ := os.Stdin.Stat()

	var output []rune

	//os.ModeCharDevice is a flag that indicates a terminal or console device.
	// info.Mode() gets the mode bits.
	// The & (bitwise AND) checks if that flag is set.
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gorainbow")
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	typewriter(output, 4*time.Millisecond)
}

func typewriter(text []rune, delay time.Duration) {
	for j, ch := range text {
		r, g, b := rgb(j) // GOLD COLOR
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, ch)
		time.Sleep(delay)
	}
	fmt.Println()
}

// rgb produce rainbow color rgb
func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

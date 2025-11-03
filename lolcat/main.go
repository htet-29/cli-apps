package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"syreclabs.com/go/faker"
)

func main() {

	var phrases []string

	for _ = range 2 {
		phrases = append(phrases, faker.Hacker().Phrases()...)
	}

	output := strings.Join(phrases[:], "; ")

	typewriter(output, 4*time.Millisecond)

}

func typewriter(text string, delay time.Duration) {
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

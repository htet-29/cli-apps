package main

import (
	"fmt"
	"strings"
	"time"

	"syreclabs.com/go/faker"
)

func main() {
	// ========== STAGE 1: System Startup ==========
	systemMessages := []struct {
		color string
		text  string
		delay time.Duration
	}{
		{"\033[36m", "Initializing J.A.R.V.I.S. OS v1.9.5...\033[0m", 1 * time.Second},
		{"\033[33m", "Loading neural modules...", 600 * time.Millisecond},
	}

	for _, msg := range systemMessages {
		fmt.Println(msg.color + msg.text + "\033[0m")
		time.Sleep(msg.delay)
	}

	// ========== STAGE 2: Loading Bar ==========
	loadingBar(30, 3*time.Second, "\033[36m")

	// Continue boot sequence
	moreMessages := []struct {
		color string
		text  string
		delay time.Duration
	}{
		{"\033[35m", "Establishing network uplink...", 900 * time.Millisecond},
		{"\033[32m", "System online. Access granted.\033[0m", 1 * time.Second},
	}

	for _, msg := range moreMessages {
		fmt.Println(msg.color + msg.text + "\033[0m")
		time.Sleep(msg.delay)
	}

	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// ========== STAGE 3: Gold Typing Effect ==========
	var phrases []string

	for _ = range 2 {
		phrases = append(phrases, faker.Hacker().Phrases()...)
	}

	output := strings.Join(phrases[:], "; ")
	r, g, b := 255, 215, 0 // GOLD COLOR

	typewriter(output, r, g, b, 4*time.Millisecond)

	// ========== STAGE 4: Blinking Cursor ==========
	for _ = range 6 {
		fmt.Print("\033[5m_\033[0m")
		time.Sleep(300 * time.Millisecond)
		fmt.Print("\b \b")
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Println("\n\033[36mAwaiting next command, Sir.\033[0m")
}

func typewriter(text string, r, g, b int, delay time.Duration) {
	for _, ch := range text {
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, ch)
		time.Sleep(delay)
	}
}

func loadingBar(width int, duration time.Duration, color string) {
	steps := width
	delay := duration / time.Duration(steps)
	for i := range steps + 1 {
		percent := (i * 100) / steps
		fmt.Printf("\r%s[%-*s] %3d%%", color, width, strings.Repeat("█", i), percent)
		time.Sleep(delay)
	}
	fmt.Print("\033[0m\n")
}

package main

import (
	"log"
	"os"

	"github.com/rthornton128/goncurses"
)

func main() {

	stdscr, err := goncurses.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer goncurses.End()
	// Configure ncurses settings
	goncurses.Echo(false)  // Disable echoing of typed characters
	goncurses.CBreak(true) // Enable character-at-a-time input

	// Print text to the screen
	stdscr.Print("Hello, ncurses in Go!")
	stdscr.Refresh()

	// Wait for user input (e.g., a key press)
	stdscr.GetChar()

	os.Exit(1)
}

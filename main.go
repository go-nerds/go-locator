package main

import (
	"log"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/gookit/color"
	"github.com/pterm/pterm"
)

type UserInput struct {
	Ip     string `arg:"--ip"`
	File   string `arg:"-f"`
	Output string `arg:"-o"`
}

func main() {

	info := color.FgLightCyan.Render

	pterm.DefaultBox.
		WithRightPadding(10).
		WithLeftPadding(10).
		WithTopPadding(2).
		WithBottomPadding(2).
		Println(info("Go locator created by AAVision"))

	userInput := UserInput{}
	arg.MustParse(&userInput)

	color.Yellowln("Processing...")
	start := time.Now()

	err := initReader(userInput)

	if err != nil {
		log.Fatal(err)
	}

	duration := time.Since(start)
	color.Cyanln("Finished in:", duration)

}

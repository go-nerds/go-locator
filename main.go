package main

import (
	"log"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/gookit/color"
)

type UserInput struct {
	Ip     string `arg:"--ip"`
	File   string `arg:"-f"`
	Output string `arg:"-o"`
}

func main() {

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

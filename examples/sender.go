package main

import (
	"fmt"
	sender2 "github.com/RickHulzinga/go-simple-artnet/sender"
	"github.com/RickHulzinga/go-simple-artnet/universe"
	"log"
	"time"
)

func main() {
	// Create a new ArtNetSender
	sender, err := sender2.NewArtNetSender("255.255.255.255:6454")
	if err != nil {
		log.Fatalf("Failed to create ArtNet sender: %v", err)
	}
	defer sender.Close()

	u, err := universe.NewDMXUniverse(0)

	value := 0
	step := 10
	increasing := true

	for {
		fmt.Printf("Value: %d\n", value)
		if increasing {
			value += step
			if value >= 255 {
				value = 255
				increasing = false
			}
		} else {
			value -= step
			if value <= 0 {
				value = 0
				increasing = true
			}
		}
		time.Sleep(100 * time.Millisecond) // Adding a small delay for better real-time

		u.SetChannel(1, value)
		sender.Send(u)
	}
}

package main

import (
	"Golang-Artnet/node"
	"time"
)

func main() {

	n, err := node.NewArtNetNode("255.255.255.255:6454")

	n.SetFrameRate(30)

	n.Start()

	if err != nil {
		panic(err)
	}

	u := n.GetUniverse(0)

	for {
		//Loop from channel 1 to 3
		for c := 1; c <= 3; c++ {

			//Fade the channel from 0 to 255
			for i := 0; i <= 255; i++ {
				u.SetChannel(c, i)
				time.Sleep(10 * time.Millisecond)
			}
			u.Blackout()

			strobeDelay := 100 * time.Millisecond

			//Now strobe
			for s := 0; s <= 30; s++ {
				u.SetChannel(c, 255)
				time.Sleep(strobeDelay)
				u.SetChannel(c, 0)
				time.Sleep(strobeDelay)
			}

		}
	}

}

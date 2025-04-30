package node

import (
	"github.com/RickHulzinga/go-simple-artnet/sender"
	"github.com/RickHulzinga/go-simple-artnet/universe"
	"time"
)

type ArtNetNode struct {
	sender    *sender.ArtNetSender
	universes map[uint16]*universe.DMXUniverse
	frameRate int

	stopChan chan struct{}
}

func (n *ArtNetNode) createUniverse(id uint16) *universe.DMXUniverse {
	universe, err := universe.NewDMXUniverse(id)
	if err != nil {
		panic(err)
	}
	return universe
}

func (n *ArtNetNode) GetUniverse(id uint16) *universe.DMXUniverse {
	// Check if the universe exists, if not, create it
	if n.universes[id] == nil {
		n.universes[id] = n.createUniverse(id)
	}
	return n.universes[id]
}

// worker is the main working function for the ArtNet node
func (n *ArtNetNode) worker() {

	frameDuration := time.Second / time.Duration(n.frameRate)
	ticker := time.NewTicker(frameDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			for _, dmxUniverse := range n.universes {
				n.sender.Send(dmxUniverse)
				time.Sleep(100 * time.Millisecond)
			}

		case <-n.stopChan: // Stop signal received
			return
		}
	}

}

// Start begins the operation of the ArtNetNode
func (n *ArtNetNode) Start() {
	n.stopChan = make(chan struct{})
	go n.worker()
}

// Stop stops the operation of the ArtNetNode
func (n *ArtNetNode) Stop() {
	// Send a signal to the stop channel
	close(n.stopChan)
}

func (n *ArtNetNode) FrameRate() int {
	return n.frameRate
}

func (n *ArtNetNode) SetFrameRate(frameRate int) {
	n.frameRate = frameRate
}

// NewArtNetNode creates a new ArtNet node which targets the specified ip
func NewArtNetNode(target string) (*ArtNetNode, error) {

	sender, err := sender.NewArtNetSender(target)

	if err != nil {
		return nil, err
	}

	return &ArtNetNode{
		sender:    sender,
		universes: make(map[uint16]*universe.DMXUniverse),
		frameRate: 30,
	}, nil
}

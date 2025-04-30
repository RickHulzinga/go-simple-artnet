package universe

import (
	"fmt"
)

type DMXUniverse struct {
	universe    uint16
	channelData []byte
}

// FullOn sets all channels to 255
func (u *DMXUniverse) FullOn() {
	for i, _ := range u.channelData {
		u.channelData[i] = 255
	}
}

// Blackout sets all channels to 0
func (u *DMXUniverse) Blackout() {
	for i, _ := range u.channelData {
		u.channelData[i] = 0
	}
}

// ChannelData byte slice with the channel values where index specifies channel
func (U *DMXUniverse) ChannelData() []byte {
	return U.channelData
}

// Universe dmx universe number
func (U *DMXUniverse) Universe() uint16 {
	return U.universe
}

func (U *DMXUniverse) GetChannel(channel int) int {
	return int(U.channelData[channel-1])
}

// SetChannel sets dmx channel to specified value
func (U *DMXUniverse) SetChannel(channel int, value int) error {
	if channel < 1 || channel > 512 {
		return fmt.Errorf("channel must be between 1 and 512")
	}
	if value < 0 {
		value = 0
	}
	if value > 255 {
		value = 255
	}

	U.channelData[channel-1] = byte(value)

	return nil
}

// NewDMXUniverse creates new universe for universe id
func NewDMXUniverse(universe uint16) (*DMXUniverse, error) {

	if universe < 0 {
		return nil, fmt.Errorf("universe cant be lower than 0")
	}

	return &DMXUniverse{
		universe:    universe,
		channelData: make([]byte, 512),
	}, nil
}

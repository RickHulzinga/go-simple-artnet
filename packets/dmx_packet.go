package packets

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/RickHulzinga/go-simple-artnet/universe"
)

type DMXPacket struct {
	Sequence uint16
	Universe uint16 // DMX Universe to target (0-15 bits allowed)
	DMXData  []byte // The DMX512 data (up to 512 channels)
}

func NewDMXPacket(universe *universe.DMXUniverse, sequence uint16) (*DMXPacket, error) {
	// Validate that the universe is not nil
	if universe == nil {
		return nil, fmt.Errorf("universe cannot be nil")
	}

	// Create and return the DMXPacket using the provided universe
	return &DMXPacket{
		Universe: universe.Universe(),
		DMXData:  universe.ChannelData(),
		Sequence: sequence,
	}, nil
}

//// NewDMXPacket creates a new packet for a single universe with DMX data
//func NewDMXPacket(universe uint16, channel int, value byte) (*DMXPacket, error) {
//	if channel < 1 || channel > 512 {
//		return nil, fmt.Errorf("channel must be between 1 and 512")
//	}
//
//	// Create DMX data with all zeros
//	dmxData := make([]byte, 512)
//
//	// Update the specific channel (convert 1-based channel to 0-based index)
//	dmxData[channel-1] = value
//
//	return &DMXPacket{
//		Universe: universe,
//		DMXData:  dmxData,
//	}, nil
//}

// ToBytes converts the Art-DMX packet into bytes for UDP transmission
func (p *DMXPacket) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Header: "Art-Net\0"
	buf.Write([]byte("Art-Net\x00"))

	// OpCode: 0x5000 for ArtDMX (Little-Endian)
	binary.Write(buf, binary.LittleEndian, uint16(0x5000))

	// Protocol Version (Hi-Lo): Currently 14
	binary.Write(buf, binary.BigEndian, uint16(14))

	// Sequence (set to 0 for simplicity), Physical (set to 0)
	buf.WriteByte(byte(p.Sequence)) // Sequence
	buf.WriteByte(0)                // Physical Port

	// Universe (Low Byte first, Little-Endian format)
	binary.Write(buf, binary.LittleEndian, p.Universe)

	// DMX Length (length of DMX data, always 512 bytes, in Big-Endian format)
	binary.Write(buf, binary.BigEndian, uint16(len(p.DMXData)))

	// DMX Data Payload
	buf.Write(p.DMXData)

	return buf.Bytes(), nil
}

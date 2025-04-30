package sender

import (
	"Golang-Artnet/packets"
	"Golang-Artnet/universe"
	"net"
)

// ArtNetSender handles sending Art-Net packets over UDP
type ArtNetSender struct {
	conn     *net.UDPConn
	target   string
	sequence uint16
}

// incrementSequence increments the sequence number, resetting it to 0 if it reaches 255.
func (a *ArtNetSender) incrementSequence() {
	if a.sequence == 255 {
		a.sequence = 0
	}
	a.sequence++
}

// NewArtNetSender creates a new ArtNet UDP sender
func NewArtNetSender(target string) (*ArtNetSender, error) {

	// No address binding required for sending packets
	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		return nil, err
	}

	return &ArtNetSender{
		conn:     conn,
		target:   target,
		sequence: 0,
	}, nil
}

// SendRaw sends raw bytes to the target IP and port
func (s *ArtNetSender) SendRaw(data []byte) error {
	// Resolve the target address (e.g., "192.168.1.255:6454" for broadcast)
	targetAddr, err := net.ResolveUDPAddr("udp", s.target)
	if err != nil {
		return err
	}

	// Send the data
	_, err = s.conn.WriteToUDP(data, targetAddr)

	s.incrementSequence()

	return err
}

// Close closes the UDP connection
func (s *ArtNetSender) Close() error {
	return s.conn.Close()
}

// Send creates a DMX packet from the given universe and sends it
func (s *ArtNetSender) Send(universe *universe.DMXUniverse) error {
	packet, err := packets.NewDMXPacket(universe, s.sequence)
	if err != nil {
		panic(err)
	}

	data, err := packet.ToBytes()
	if err != nil {
		panic(err)
	}

	return s.SendRaw(data)
}

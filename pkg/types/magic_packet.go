// The magic packet is a frame that is most often sent as a broadcast
// and that contains anywhere within its payload 6 bytes of all 255 (FF FF FF FF FF FF in hexadecimal),
// followed by sixteen repetitions of the target computer's 48-bit MAC address, for a total of 102 bytes.
//
// https://en.wikipedia.org/wiki/Wake-on-LAN#Magic_packet

package types

func NewMagicPacket(mac string) (*[]byte, error) {
	var packet []byte

	// 6 repetitions of 0xFF
	for i := 0; i < 6; i++ {
		packet = append(packet, 0xFF)
	}

	m, err := NewMacAddress(mac)
	if err != nil {
		return nil, err
	}

	// 16 repetitions of target MAC
	b, err := m.Bytes()
	if err != nil {
		return nil, err
	}
	for i := 0; i < 16; i++ {
		packet = append(packet, b...)
	}

	return &packet, nil
}

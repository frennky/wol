package wol

import (
	"fmt"
	"net"
)

const (
	DefaultBroadcastIPAddress = "255.255.255.255"
	// Magic pocket is usually sent to port 0, 7 or 9
	DefaultPort = 0
)

// Sends bytes over udp to provided broadcast address
func Send(host string, port int, packet *[]byte) error {
	bcastAddr := fmt.Sprintf("%s:%d", host, port)
	udpAddr, err := net.ResolveUDPAddr("udp4", bcastAddr)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	n, err := conn.Write(*packet)
	if err == nil && n != 102 {
		return fmt.Errorf("sent %d bytes, expected 102 bytes", n)
	}
	if err != nil {
		return err
	}

	return nil
}

package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"regexp"
)

const (
	MacPattern = "^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"
)

type MacAddress [6]byte

func NewMacAddress(mac string) (*MacAddress, error) {
	var macAddr MacAddress

	// Only support 6 byte MAC addresses
	if m, _ := regexp.MatchString(MacPattern, mac); !m {
		return nil, fmt.Errorf("%s is not a IEEE 802 MAC-48 address", mac)
	}

	hwAddr, err := net.ParseMAC(mac)
	if err != nil {
		return nil, err
	}

	// Copy bytes from the returned HardwareAddr
	for i := range macAddr {
		macAddr[i] = hwAddr[i]
	}

	return &macAddr, nil
}

func (ma *MacAddress) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, ma); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

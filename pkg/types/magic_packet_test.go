package types

import (
	"bytes"
	"testing"
)

func TestNewMagicPacket(t *testing.T) {
	mac := "aa:bb:cc:dd:ff:11"
	expected := &[]byte{
		255, 255, 255, 255, 255, 255,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17,
		170, 187, 204, 221, 255, 17}

	mp, err := NewMagicPacket(mac)

	if err != nil {
		t.Errorf("Expected nil, got %v", err.Error())
	}

	if !bytes.Equal(*mp, *expected) {
		t.Errorf("Expected %v, got %v", expected, mp)
	}
}

func TestNewMagicPacketWithInvalidMac(t *testing.T) {
	mac := "abc"

	mp, err := NewMagicPacket(mac)

	if mp != nil {
		t.Errorf("Expected nil, got %v", mp)
	}

	if err == nil {
		t.Error("Expected error got nil.")
	}
}

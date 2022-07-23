package types

import (
	"bytes"
	"testing"
)

func TestNewMacAddress(t *testing.T) {
	mac := "aa:bb:cc:dd:ff:11"

	_, err := NewMacAddress(mac)

	if err != nil {
		t.Errorf("Expected nil, got %v", err.Error())
	}
}

func TestNewMacAddressWithInvalidMac(t *testing.T) {
	macs := []string{
		"abc",
		"aa:bb:cc:dd:ff:11 aa:bb:cc:dd:ff:11",
		"zz:zz:zz:zz:zz:zz",
	}

	for _, mac := range macs {

		ma, err := NewMacAddress(mac)

		if ma != nil {
			t.Errorf("Expected nil, got %v", ma)
		}

		if err == nil {
			t.Error("Expected error, got nil.")
		}

	}
}

func TestBytes(t *testing.T) {
	mac := "aa:bb:cc:dd:ff:11"

	ma, err := NewMacAddress(mac)

	if err != nil {
		t.Errorf("Expected nil, got %v", err.Error())
	}

	ab, err := ma.Bytes()

	if err != nil {
		t.Errorf("Expected nil, got %v", err.Error())
	}

	eb := []byte{170, 187, 204, 221, 255, 17}

	if !bytes.Equal(ab, eb) {
		t.Errorf("Expected: %v, got %v", eb, ab)
	}
}

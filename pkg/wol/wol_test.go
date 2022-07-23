package wol

import (
	"testing"
)

func TestSend(t *testing.T) {
	b := []byte{}
	for i := 0; i < 102; i++ {
		b = append(b, byte(i))
	}

	err := Send(DefaultBroadcastIPAddress, DefaultPort, &b)

	if err != nil {
		t.Errorf("Expected nil, got %v", err.Error())
	}
}

func TestSendWithInvalidPacketSize(t *testing.T) {
	b := []byte{}

	err := Send(DefaultBroadcastIPAddress, DefaultPort, &b)

	if err == nil {
		t.Error("Expected error, got nil.")
	}
	for i := 0; i < 103; i++ {
		b = append(b, byte(i))
	}

	err = Send(DefaultBroadcastIPAddress, DefaultPort, &b)

	if err == nil {
		t.Error("Expected error, got nil.")
	}
}

func TestSendWithInvalidHost(t *testing.T) {
	b := []byte{}
	for i := 0; i < 102; i++ {
		b = append(b, byte(i))
	}

	err := Send("299.299.299.299", DefaultPort, &b)

	if err == nil {
		t.Error("Expected error, got nil.")
	}
}

func TestSendWithInvalidPort(t *testing.T) {
	b := []byte{}
	for i := 0; i < 102; i++ {
		b = append(b, byte(i))
	}

	err := Send(DefaultBroadcastIPAddress, 99999, &b)

	if err == nil {
		t.Error("Expected error, got nil.")
	}
}

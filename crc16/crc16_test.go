package crc16

import (
	"testing"
)

func TestCompute(t *testing.T) {
	data := []byte("123456789")
	crc := Compute(data)
	if crc != 0x31C3 {
		t.Errorf("Expected CRC16 of '123456789' to be 0x31C3, got %X", crc)
	}
}

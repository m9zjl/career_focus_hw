package utils

import (
	"testing"
)

func TestHexToInt(t *testing.T) {
	talbes := []struct {
		hexValue string
		intValue int64
	}{{"0xa", 10}}
	for _, table := range talbes {
		ret := HexToInt(table.hexValue)
		if ret != table.intValue {
			t.Errorf("HexToInt(%s) returned %d, want %d", table.hexValue, ret, table.intValue)
		}

	}
}

func TestIntToHnt(t *testing.T) {
	talbes := []struct {
		hexValue string
		intValue int64
	}{{"0xa", 10}}
	for _, table := range talbes {
		ret := IntToHex(table.intValue)
		if ret != table.hexValue {
			t.Errorf("IntToHex(%d) returned %s, want %s", table.intValue, ret, table.hexValue)
		}

	}
}

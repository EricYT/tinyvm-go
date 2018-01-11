package tvm

import (
	"testing"
)

func TestParseValue(t *testing.T) {
	val := []byte("123")
	if ParseValue(val) != 123 {
		t.Fatalf("TestParseValue parse %s got wrong value %d", string(val), ParseValue(val))
	}

	val = []byte("0x1a")
	if ParseValue(val) != 26 {
		t.Fatalf("TestParseValue parse %s got wrong value %d", string(val), ParseValue(val))
	}

	val = []byte("01010")
	if ParseValue(val) != 520 {
		t.Fatalf("TestParseValue parse %s got wrong value %d", string(val), ParseValue(val))
	}

	val = []byte("1a30|h")
	if ParseValue(val) != 6704 {
		t.Fatalf("TestParseValue parse %s got wrong value %d", string(val), ParseValue(val))
	}

	val = []byte("1010|b")
	if ParseValue(val) != 10 {
		t.Fatalf("TestParseValue parse %s got wrong value %d", string(val), ParseValue(val))
	}
}

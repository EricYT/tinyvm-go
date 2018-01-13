package tvm

import "testing"

func TestOpCodeMap(t *testing.T) {
	if len(opcodeMap) != len(opcodes) {
		t.Fatalf("The length of opcodeMap and opcodes is not equal")
	}

	for opcode, opcodeName := range opcodes {
		c, ok := opcodeMap[opcodeName]
		if !ok {
			t.Fatalf("Opcode %s not found in opcodeMap", opcodeName)
		}
		if Opcode(opcode) != c {
			t.Fatalf("Opcode %s code(%d) is not same as opcodeMap(%d)", opcodeName, opcode, c)
		}
	}
}

func TestRegisterMap(t *testing.T) {
	if len(registerMap) != len(registers) {
		t.Fatalf("The length of registerMap and registers is not equal")
	}

	for regIdx, regName := range registers {
		reg, ok := registerMap[regName]
		if !ok {
			t.Fatalf("Register %s is not found in registerMap", regName)
		}
		if Register(regIdx) != reg {
			t.Fatalf("Register %s register index(%d) is not same as registerMap(%d)", regName, regIdx, reg)
		}
	}
}

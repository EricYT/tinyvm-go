package tvm

import "fmt"

// tvm memory

const (
	NUM_REGISTERS int = 17

	MIN_MEMORY_SIZE int = 64 * 1024 * 1024 /* 64 MB */
)

type Mem struct {
	/*
	 * Similar to x86 FLAGS register
	 *
	 * 0x1 EQUAL
	 * 0x2 GREATER
	 *
	 */

	FLAGS     int
	remainder int

	space     []int
	spaceSize int

	registers [17]regUnit
}

func NewMem(size int) *Mem {
	m := new(Mem)

	m.spaceSize = size
	m.space = make([]int, size)

	return m
}

type regUnit struct {
	//i32 int32
	//val []int32
	i32    int
	i32Ptr *int
}

func (ru regUnit) String() string {
	if ru.i32Ptr != nil {
		return fmt.Sprintf("i32: %d i32Ptr: %d", ru.i32, *ru.i32Ptr)
	}
	return fmt.Sprintf("i32: %d i32Ptr is nil", ru.i32)
}

/* Initialize our stack by setting the base pointer and stack pointer */
func (m *Mem) StackCreate(size int) {
	// point to the bottom of stack
	m.registers[EBP].i32 = len(m.space)
	m.registers[ESP].i32 = m.registers[EBP].i32
}

func (m *Mem) StackPush(item *int) {
	m.registers[ESP].i32 -= 1
	m.space[m.registers[ESP].i32] = *item
}

func (m *Mem) StackPop() *int {
	dest := m.space[m.registers[ESP].i32]
	m.registers[ESP].i32 += 1
	return &dest
}

func (m *Mem) SetRegisterI32(reg Register, v int) {
	m.registers[reg].i32 = v
}

func (m *Mem) SetRegisterI32Ptr(reg Register, v *int) {
	m.registers[reg].i32Ptr = v
}

func (m *Mem) GetRegisterI32(reg Register) *int {
	return &m.registers[reg].i32
}

func (m *Mem) GetRegisterI32Ptr(reg Register) *int {
	return m.registers[reg].i32Ptr
}

func (m *Mem) GetFLAGS() int { return m.FLAGS }

func (m *Mem) GetRemainder() int { return m.remainder }

func (m *Mem) AllRegisterInfos() string {
	var infos string
	for idx, regName := range registers {
		regInfo := fmt.Sprintf("register %s value: %s\n", regName, m.registers[idx])
		infos += regInfo
	}
	return infos
}

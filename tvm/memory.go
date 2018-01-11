package tvm

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

/* Initialize our stack by setting the base pointer and stack pointer */
func (m *Mem) StackCreate(size int) {
	// point to the bottom of stack
	m.registers[0x7].i32 = len(m.space)
	m.registers[0x6].i32 = m.registers[0x7].i32
}

func (m *Mem) StackPush(item *int) {
	m.registers[0x6].i32 -= 1
	m.space[m.registers[0x6].i32] = *item
}

func (m *Mem) StackPop() *int {
	dest := m.space[m.registers[0x6].i32]
	m.registers[0x6].i32 += 1
	return &dest
}

func (m *Mem) SetRegisterI32(idx int, v int) {
	m.registers[idx].i32 = v
}

func (m *Mem) SetRegisterI32Ptr(idx int, v *int) {
	m.registers[idx].i32Ptr = v
}

func (m *Mem) GetRegisterI32(idx int) *int {
	return &m.registers[idx].i32
}

func (m *Mem) GetRegisterI32Ptr(idx int) *int {
	return m.registers[idx].i32Ptr
}

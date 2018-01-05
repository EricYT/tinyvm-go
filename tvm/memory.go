package tvm

// tvm memory

const (
	NUM_REGISTERS int = 17
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
	i32 int
	val []int
}

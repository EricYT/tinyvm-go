package tvm

// tvm program

type Prog struct {
	start int

	numInstr int
	instr    []opcode
	args     [][]*int

	values    []int
	numValues int

	defines *htabCtx
	labels  *htabCtx
}

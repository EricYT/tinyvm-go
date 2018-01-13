package tvm

// tvm program

type Prog struct {
	start int

	numInstr int
	instr    []Opcode
	args     [][]*int
	tokens   [][][]byte

	values    []int
	numValues int

	defines *htabCtx
	labels  *htabCtx
}

func NewProg() *Prog {
	prog := new(Prog)

	prog.defines = NewHtabCtx()
	prog.labels = NewHtabCtx()

	return prog
}

func (prog *Prog) Start() int {
	return prog.start
}

func (prog *Prog) NumInstr() int {
	return prog.numInstr
}

func (prog *Prog) OpCode(instrIdx int) Opcode {
	return prog.instr[instrIdx]
}

func (prog *Prog) Args(instrIdx int) []*int {
	return prog.args[instrIdx]
}

func (prog *Prog) Tokens(instrIdx int) [][]byte {
	return prog.tokens[instrIdx]
}

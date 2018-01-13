package tvm

import (
	"bytes"
	"errors"
	"strconv"
)

var (
	ErrorLabelParseLabelDefineTwice error = errors.New("parser: label defined twice")

	ErrorProgParseInstrNotExists error = errors.New("parser: no instruction exists")
)

func (ctx *Ctx) ParseProgram(lexer *lexerCtx) error {

	for _, lineTokens := range lexer.Tokens() {
		opc, instrIdx := ctx.parseInstr(lineTokens)
		if opc == Opcode(-1) {
			continue
		}
		ctx.Prog.instr = append(ctx.Prog.instr, opc)
		ctx.Prog.tokens = append(ctx.Prog.tokens, lineTokens)
		args := ctx.parseArgs(lineTokens, instrIdx)
		ctx.Prog.args = append(ctx.Prog.args, args)
	}

	return nil
}

func (ctx *Ctx) ParseLabels(lexer *lexerCtx) error {
	var numInstr int
	prog := ctx.Prog
	for _, tokens := range lexer.Tokens() {
		// tokens of every signal line
		validInstruction := false
		for _, token := range tokens {
			/* The token shouldn't be empty. If it is empty, or non-existent, skip it. */
			if len(token) == 0 {
				continue
			}

			/* check the token for a valid instruction */
			if _, ok := instrToOpCode(token); ok {
				validInstruction = true
			}

			/* Check for a label delimiter */
			if token[len(token)-1] != ':' {
				continue
			}
			// trim the last char ':'
			token = token[:len(token)-1]

			/* If the label is "start", make it the entry point */
			if "start" == string(token) {
				prog.start = numInstr
			}

			/* Check if the label already exists */
			if _, ok := prog.labels.Find(string(token)); ok {
				return ErrorLabelParseLabelDefineTwice
			}

			prog.labels.Add(string(token), numInstr)
		}
		if validInstruction {
			numInstr++
		}
	}

	return nil
}

/* This function takes the instruction tokens, and location of the
 * instruction inside the line, parses the arguments, and returns a
 * pointer to the heap where they're stored.
 */
func (ctx *Ctx) parseArgs(instrTokens [][]byte, index int) []*int {
	var args []*int
	for i := 0; i < MAX_ARGS; i++ {
		argIdx := index + 1 + i
		if argIdx >= len(instrTokens) {
			continue
		}

		token := instrTokens[argIdx]
		if len(token) == 0 {
			continue
		}

		/* Check to see if the token specifies a register */
		if reg := tokenToRegister(token, ctx.Mem); reg != nil {
			args = append(args, reg)
			continue
		}

		/* Check to see wheather the token specifies an address */
		if token[0] == '[' {
			endIdx := bytes.IndexByte(token, ']')
			if endIdx != -1 {
				args = append(args, &ctx.Mem.space[ParseValue(token[1:endIdx])])
				continue
			}
		}

		/* Check if the argument is a label */
		if addr, ok := ctx.Prog.labels.Find(string(token)); ok {
			args = append(args, addValue(*addr, ctx.Prog))
			continue
		}

		/* Fuck it, parse it as a value */
		args = append(args, addValue(ParseValue(token), ctx.Prog))
	}
	return args
}

/*
 This is a helper function that converts one instruction,
 from one line of code, to tvm bytecode.
*/
func (ctx *Ctx) parseInstr(instrTokens [][]byte) (Opcode, int) {
	for index, token := range instrTokens {
		// skip empty one
		if len(token) == 0 {
			continue
		}
		opc, ok := instrToOpCode(token)
		if ok {
			ctx.Prog.numInstr++
			return opc, index
		}
	}

	return Opcode(-1), -1
}

// utils
func ParseValue(token []byte) int {
	delimiter := bytes.IndexByte(token, '|')
	var base int = 0

	if delimiter != -1 {
		identifier := delimiter + 1

		switch token[identifier] {
		case 'h':
			base = 16
		case 'b':
			base = 2
		default:
			base = 0
		}
		token = token[:delimiter]
	}

	val, err := strconv.ParseInt(string(token), base, 64)
	if err != nil {
		panic(err)
	}

	return int(val)
}

func instrToOpCode(instr []byte) (Opcode, bool) {
	if opc, ok := opcodeMap[string(instr)]; ok {
		return opc, true
	}
	return Opcode(-1), false
}

func tokenToRegister(token []byte, mem *Mem) *int {
	if reg, ok := registerMap[string(token)]; ok {
		return &mem.registers[reg].i32
	}
	return nil
}

func addValue(val int, prog *Prog) *int {
	prog.values = append(prog.values, val)
	return &prog.values[len(prog.values)-1]
}

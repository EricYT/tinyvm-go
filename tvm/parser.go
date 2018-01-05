package tvm

import (
	"bytes"
	"errors"
)

var (
	ErrorLabelParseLabelDefineTwice error = errors.New("parser: label defined twice")

	ErrorProgParseInstrNotExists error = errors.New("parser: no instruction exists")
)

func (ctx *Ctx) ParseLabels(lexer *lexerCtx) error {
	var numInstr int
	prog := ctx.prog

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
			token = bytes.TrimLeft(token, ":")

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
		if argIdx > len(instrTokens) || len(instrTokens[argIdx]) == 0 {
			continue
		}

		token := instrTokens[argIdx]
		/* Check to see if the token specifies a register */
		if reg := tokenToRegister(token, ctx.mem); reg != nil {
			args[i] = reg
		}
	}

	return args
}

/*
 This is a helper function that converts one instruction,
 from one line of code, to tvm bytecode.
*/
func (ctx *Ctx) parseInstr(instrTokens [][]byte) (opcode, int) {
	for index, token := range instrTokens {
		// skip empty one
		if len(token) == 0 {
			continue
		}
		if opc, ok := instrToOpCode(token); ok {
			ctx.prog.numInstr++
			return opc, index
		}
	}

	return opcode(-1), -1
}

// utils
func instrToOpCode(instr []byte) (opcode, bool) {
	if opc, ok := opcodeMap[string(instr)]; ok {
		return opc, true
	}
	return opcode(-1), false
}

func tokenToRegister(token []byte, mem *Mem) *int {
	if reg, ok := registerMap[string(token)]; ok {
		return &mem.registers[reg].i32
	}
	return nil
}

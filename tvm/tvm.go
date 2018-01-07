package tvm

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Package tvm is the core package implement the tiny virtual machine.

/* core struct tvm */

type Ctx struct {
	prog *Prog
	mem  *Mem
}

func NewCtx() *Ctx {
	ctx := new(Ctx)

	ctx.prog = NewProg()
	ctx.mem = NewMem(MIN_MEMORY_SIZE)
	ctx.mem.StackCreate(MIN_MEMORY_SIZE)

	return ctx
}

func (ctx *Ctx) Interpret(filename string) error {
	log.Printf("Prepare to interpret the file %s", filename)

	// Attempt to read the source file
	source, err := ReadFile(filename, ".vm")
	if err != nil {
		return err
	}

	// preprocess source
	if source, err = ctx.preprocess(source); err != nil {
		return err
	}

	// lexer analysis
	lexer := NewLexerCtx()
	lexer.Parse(source, ctx.prog.defines)

	// parse labels
	if err := ctx.ParseLabels(lexer); err != nil {
		return err
	}

	// parse program
	if err := ctx.ParseProgram(lexer); err != nil {
		return err
	}

	return nil
}

func (ctx *Ctx) Run() error {
	instrIdx := &ctx.mem.registers[0x8].i32
	*instrIdx = ctx.prog.start

	// program run step by step
	for {
		if *instrIdx > len(ctx.prog.instr)-1 {
			break
		}
		ctx.step(instrIdx)
		(*instrIdx)++
	}

	return nil
}

func (ctx *Ctx) step(instrIdx *int) error {
	opcode := ctx.prog.instr[*instrIdx]
	args := ctx.prog.args[*instrIdx]

	switch opcode {
	case NOP: // nop
	case INT: // int  TODO: not implement
	case MOV: // move
		*args[0] = *args[1]

	// stack operations
	case PUSH: // push
		ctx.mem.StackPush(args[0])
	case POP: // pop
		*args[0] = *ctx.mem.StackPop()
	case PUSHF: // pushf
		ctx.mem.StackPush(&ctx.mem.FLAGS)
	case POPF: // popf
		*args[0] = *ctx.mem.StackPop()

	// arithmetic operators
	case INC: // inc
		*args[0]++
	case DEC: // dec
		*args[0]--
	case ADD: // add
		*args[0] += *args[1]
	case SUB: // sub
		*args[0] -= *args[1]
	case MUL: // mul
		*args[0] *= *args[1]
	case DIV: // div
		*args[0] /= *args[1]
	case MOD: // mod
		ctx.mem.remainder = *args[0] % *args[1]
	case REM: // rem
		*args[0] = ctx.mem.remainder

	// arithmetic shift
	case NOT: // not
		*args[0] = ^(*args[0])
	case XOR: // xor
		*args[0] ^= *args[1]
	case OR: // or
		*args[0] |= *args[1]
	case AND: // and
		*args[0] &= *args[1]
	case SHL: // shl
		*args[0] <<= uint(*args[1])
	case SHR: // shr
		*args[0] >>= uint(*args[1])

	case CMP: // cmp
		var r1 int
		if *args[0] == *args[1] {
			r1 = 0x1
		}
		var r2 int
		if *args[0] > *args[1] {
			r2 = 0x1
		}
		ctx.mem.FLAGS = r1 | (r2 << 1)
	case CALL: // call
		ctx.mem.StackPush(instrIdx)
		fallthrough
	case JMP: // jmp
		*instrIdx = *args[0] - 1
	case RET: // ret
		*instrIdx = *ctx.mem.StackPop()

	case JE: // je
		if (ctx.mem.FLAGS & 0x1) == 0x1 {
			*instrIdx = *args[0] - 1
		}
	case JNE: // jne
		if (ctx.mem.FLAGS & 0x1) != 0x1 {
			*instrIdx = *args[0] - 1
		}
	case JG: // jg
		if (ctx.mem.FLAGS & 0x2) == 0x2 {
			*instrIdx = *args[0] - 1
		}
	case JGE: // jge
		if (ctx.mem.FLAGS & 0x3) != 0x0 {
			*instrIdx = *args[0] - 1
		}
	case JL: // jl
		if (ctx.mem.FLAGS & 0x3) == 0x0 {
			*instrIdx = *args[0] - 1
		}
	case JLE: // jle
		if (ctx.mem.FLAGS & 0x2) == 0x0 {
			*instrIdx = *args[0] - 1
		}
	case PRN: // prn
		fmt.Printf("%d\n", *args[0])
	}

	return nil
}

// utils
func ReadFile(filename string, extension string) ([]byte, error) {
	var source []byte
	var err error

	// read file by filename
	if source, err = ioutil.ReadFile(filename); err == nil {
		return source, nil
	}
	if strings.HasSuffix(filename, extension) {
		return nil, err
	}
	// try to read file with extension
	if source, err = ioutil.ReadFile(filename + extension); err != nil {
		return nil, err
	}
	return source, nil
}

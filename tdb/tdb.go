package tdb

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode"

	tinyvm "github.com/EricYT/tinyvm-go/tvm"
)

// a debug tool for tinyvm-go, like gdb to c, c++

func Shell(tvm *tinyvm.Ctx) error {
	var running bool = false
	var breakpoints []BreakPoint

	// initialize start register
	tvm.Mem.SetRegisterI32(tinyvm.EIP, tvm.Prog.Start())

	// here we go
	reader := bufio.NewReader(os.Stdin)

LOOP:
	for {
		fmt.Printf("tdb >: ")
		// read input
		input, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err != nil && err == io.EOF {
			break
		}

		// operate args
		tokens := tokenize(input)
		if len(tokens) == 0 {
			fmt.Printf("WARNING: no input.\n")
			continue
		}
		cmd, _ := inputToCmd(string(tokens[0]))

		// cmd operation
		var hitBreakPoint bool
		switch cmd {
		default:
			fmt.Printf("WARNING: \"%s\" is not a valid command.\n", string(tokens[0]))
		case CMD_QUIT:
			break LOOP
		case CMD_RUN:
			if running {
				fmt.Printf("The program is alrady running.\n")
				continue
			}
			hitBreakPoint = run(tvm, breakpoints)
			running = true
		case CMD_BREAK:
			addr := tinyvm.ParseValue(tokens[1])
			bp := BreakPoint{address: addr}
			breakpoints = append(breakpoints, bp)
		case CMD_STEP:
			tvm.Step(tvm.Mem.GetRegisterI32(tinyvm.EIP))
			(*tvm.Mem.GetRegisterI32(tinyvm.EIP))++
			fmt.Printf("Advancing instruction pointer to %d\n", *tvm.Mem.GetRegisterI32(tinyvm.EIP))
		case CMD_CONTINUE:
			// one step jumps over break point.
			tvm.Step(tvm.Mem.GetRegisterI32(tinyvm.EIP))
			(*tvm.Mem.GetRegisterI32(tinyvm.EIP))++
			hitBreakPoint = run(tvm, breakpoints)
		case CMD_INFOS:
			displayInfos(tvm)
		case CMD_ARGS:
			displayArgs(tvm)
		case CMD_INSTR:
			displayInstruction(tvm)
		}

		if tvm.End() {
			fmt.Printf("End of program readched.\n")
			break LOOP
		}

		if hitBreakPoint {
			fmt.Printf("BreakPoint hit at address: %d\n", *tvm.Mem.GetRegisterI32(tinyvm.EIP))
		}
	}

	return nil
}

func displayInfos(tvm *tinyvm.Ctx) {
	regInfos := tvm.Mem.AllRegisterInfos()
	fmt.Printf("register infos: \n%s", regInfos)
	fmt.Printf("FLAGS: %d Remainder: %d\n", tvm.Mem.GetFLAGS(), tvm.Mem.GetRemainder())
}

func displayArgs(tvm *tinyvm.Ctx) {
	eip := *tvm.Mem.GetRegisterI32(tinyvm.EIP)
	args := tvm.Prog.Args(eip)
	if args != nil {
		opcode := tvm.Prog.OpCode(eip)
		fmt.Printf("current program eip: %d opcode: %d instruction: %s\nargs: ", eip, opcode, opcode)
		for _, arg := range args {
			fmt.Printf("%d\t", *arg)
		}
	}
	fmt.Printf("\n")
}

func displayInstruction(tvm *tinyvm.Ctx) {
	eip := *tvm.Mem.GetRegisterI32(tinyvm.EIP)
	tokens := tvm.Prog.Tokens(eip)
	if tokens != nil {
		fmt.Printf("Instruction index %d: ", eip)
		for _, token := range tokens {
			fmt.Printf("%s ", string(token))
		}
		fmt.Printf("\n")
	}
}

func run(tvm *tinyvm.Ctx, bps []BreakPoint) bool {
	var idxInstr *int = tvm.Mem.GetRegisterI32(tinyvm.EIP)
	for {
		if tvm.End() {
			return false
		}
		// check break points
		for _, bp := range bps {
			if bp.address == *idxInstr {
				// breakpoint hit
				return true
			}
		}
		// program run one step
		tvm.Step(idxInstr)
		// next step
		(*idxInstr)++
	}
	return false
}

func tokenize(input []byte) [][]byte {
	return bytes.FieldsFunc(input, func(c rune) bool {
		return unicode.IsSpace(c)
	})
}

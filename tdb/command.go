package tdb

type cmd int

const (
	CMD_QUIT cmd = iota
	CMD_RUN
	CMD_BREAK
	CMD_STEP
	CMD_CONTINUE
	CMD_INFOS
	CMD_ARGS
	CMD_INSTR
	CMD_NOP
)

var cmdMap map[string]cmd = map[string]cmd{
	"q":        CMD_QUIT,
	"run":      CMD_RUN,
	"r":        CMD_RUN,
	"break":    CMD_BREAK,
	"b":        CMD_BREAK,
	"step":     CMD_STEP,
	"s":        CMD_STEP,
	"continue": CMD_CONTINUE,
	"c":        CMD_CONTINUE,
	"infos":    CMD_INFOS,
	"i":        CMD_INFOS,
	"args":     CMD_ARGS,
	"a":        CMD_ARGS,
	"instr":    CMD_INSTR,
}

func inputToCmd(input string) (cmd, bool) {
	if c, ok := cmdMap[input]; ok {
		return c, true
	}
	return CMD_NOP, false
}

func cmdValidate(i string) (cmd, bool) {
	if c, ok := cmdMap[i]; ok {
		return c, true
	}
	return CMD_NOP, false
}

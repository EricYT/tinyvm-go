package tdb

type cmd int

const (
	CMD_QUIT cmd = iota
	CMD_RUN
	CMD_BREAK
	CMD_STEP
	CMD_CONTINUE
	CMD_NOP
)

var cmdMap map[string]cmd = map[string]cmd{
	"q":        CMD_QUIT,
	"run":      CMD_RUN,
	"break":    CMD_BREAK,
	"step":     CMD_STEP,
	"continue": CMD_CONTINUE,
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

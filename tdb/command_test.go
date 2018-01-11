package tdb

import "testing"

func TestInputToCmd(t *testing.T) {
	for cmdName, cmd := range cmdMap {
		if c, ok := inputToCmd(cmdName); !ok || c != cmd {
			t.Fatalf("cmdName: %s cmd: %d convert can't be wrong cmd: %d exists: %t", cmdName, cmd, c, ok)
		}
	}

	if _, ok := inputToCmd("foo"); ok {
		t.Fatalf("cmdName: foo shouldn't exists.")
	}
}

func TestCmdValidate(t *testing.T) {
	for cmdName, cmd := range cmdMap {
		if c, ok := cmdValidate(cmdName); !ok || c != cmd {
			t.Fatalf("cmdName: %s cmd: %d should be valid error cmd: %d exists: %t", cmdName, cmd, c, ok)
		}
	}

	if _, ok := cmdValidate("foo"); ok {
		t.Fatalf("cmdName: foo shouldn't be valid.")
	}
}

package tvm

import "testing"

func TestTvm(t *testing.T) {
	tvm := NewCtx()

	if err := tvm.Interpret("./preprocessor/jsr.vm"); err != nil {
		t.Fatalf("TestTvm source interpret error: %s", err)
	}

	tvm.Run()
}

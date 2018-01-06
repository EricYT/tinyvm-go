package tvm

import "testing"

func TestTvm(t *testing.T) {
	tvm := NewCtx()

	if err := tvm.Interpret("./preprocessor/jsr.vm"); err != nil {
		t.Fatalf("TestTvm source interpret error: %s", err)
	}

	tvm.Run()
}

func TestTvmFact(t *testing.T) {
	tvm := NewCtx()

	if err := tvm.Interpret("./fact.vm"); err != nil {
		t.Fatalf("TestTvmFact source interpret error: %s", err)
	}

	tvm.Run()
}

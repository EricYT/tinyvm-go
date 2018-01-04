package tvm

import (
	"bytes"
	"testing"
)

func TestIsComment(t *testing.T) {
	cmt := []byte("foo")
	if isComment(cmt) {
		t.Fatalf("TestIsComment (%s) is not comment", string(cmt))
	}

	cmt = []byte("# foo")
	if !isComment(cmt) {
		t.Fatalf("TestIsComment (%s) is comment", string(cmt))
	}

	cmt = []byte("  # foo")
	if !isComment(cmt) {
		t.Fatalf("TestIsComment (%s) is comment", string(cmt))
	}

	cmt = []byte(" \t \t # foo")
	if !isComment(cmt) {
		t.Fatalf("TestIsComment (%s) is comment", string(cmt))
	}

	cmt = []byte("")
	if isComment(cmt) {
		t.Fatalf("TestIsComment empty slice is not comment")
	}
}

func TestTrimSpaceLeft(t *testing.T) {
	val := []byte("foo")
	if !bytes.Equal(val, trimSpaceLeft(val)) {
		t.Fatalf("TestTrimSpaceLeft (%s) not contains space or tab")
	}

	val = []byte("    foo")
	if !bytes.Equal([]byte("foo"), trimSpaceLeft(val)) {
		t.Fatalf("TestTrimSpaceLeft (%s) should be (foo) not (%s)", string(val), string(trimSpaceLeft(val)))
	}

	val = []byte("  \t \t    foo")
	if !bytes.Equal([]byte("foo"), trimSpaceLeft(val)) {
		t.Fatalf("TestTrimSpaceLeft (%s) should be (foo) not (%s)", string(val), string(trimSpaceLeft(val)))
	}

}

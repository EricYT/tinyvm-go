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

func TestIsSpaceLine(t *testing.T) {
	val := []byte("xxx	\t dsfd\n")
	if isSpaceLine(val) {
		t.Fatalf("TestIsSpaceLine val(%s) is not space line.", string(val))
	}
	val = []byte(" \t	\n\n\t 	\t \n")
	if !isSpaceLine(val) {
		t.Fatalf("TestIsSpaceLine val(%v) is not space line.", val)
	}
}

var src string = `
	# I'm a comment.
start:
  mov eax, 42
	mov ebx, ONE
	call print_eax

	mov		eax,					23
	mov ebx,	ZERO
	call		  print_eax

`

var tokens [][][]byte = [][][]byte{
	[][]byte{
		[]byte("start:"),
	},
	[][]byte{
		[]byte("mov"),
		[]byte("eax"),
		[]byte("42"),
	},
	[][]byte{
		[]byte("mov"),
		[]byte("ebx"),
		[]byte("1"),
	},
	[][]byte{
		[]byte("call"),
		[]byte("print_eax"),
	},
	[][]byte{
		[]byte("mov"),
		[]byte("eax"),
		[]byte("23"),
	},
	[][]byte{
		[]byte("mov"),
		[]byte("ebx"),
		[]byte("0"),
	},
	[][]byte{
		[]byte("call"),
		[]byte("print_eax"),
	},
}

func TestLexerParse(t *testing.T) {
	def := NewHtabCtx()
	def.AddRef("ONE", []byte("1"))
	def.AddRef("ZERO", []byte("0"))

	lexer := NewLexerCtx()
	lexer.Parse([]byte(src), def)

	for lineIndex, lineTokens := range lexer.Tokens() {
		for tokenIndex, token := range lineTokens {
			rightToken := tokens[lineIndex][tokenIndex]
			if !bytes.Equal(token, rightToken) {
				t.Fatalf("TestLexerParse lexer parse line(%d) token(%d) value(%#v) got wrong tokens(%#v).", lineIndex, tokenIndex, rightToken, token)
			}
		}
	}
}

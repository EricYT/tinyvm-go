package tvm

import (
	"bytes"
	"unicode"
)

// lexer implement

type Token []byte
type Line []byte

type lexerCtx struct {
	sourceLines [][]byte
	tokens      [][][]byte
}

func NewLexerCtx() *lexerCtx {
	lexer := new(lexerCtx)
	return lexer
}

func (lexer *lexerCtx) Parse(src []byte, defines *htabCtx) {
	// split all source into lines
	lexer.sourceLines = bytes.Split(src, []byte("\n"))

	lexer.tokens = make([][][]byte, len(lexer.sourceLines))

	// split all lines into tokens
	for lineIndex, line := range lexer.sourceLines {
		// filter comments delimited by '#'
		if isComment(line) {
			continue
		}
		// split lines
		toks := bytes.FieldsFunc(line, func(c rune) bool {
			return c == ' ' || c == '\t' || c == ','
		})
		lexer.tokens[lineIndex] = make([][]byte, len(toks))
		for tokIndex, tok := range toks {
			if value, ok := defines.FindRef(string(tok)); ok {
				lexer.tokens[lineIndex][tokIndex] = value
				continue
			}
			lexer.tokens[lineIndex][tokIndex] = tok
		}
	}
}

func isComment(line []byte) bool {
	clear := trimSpaceLeft(line)
	if len(clear) == 0 {
		return false
	}
	return clear[0] == '#'
}

func trimSpaceLeft(value []byte) []byte {
	return bytes.TrimLeftFunc(value, func(c rune) bool {
		return unicode.IsSpace(c)
	})
}

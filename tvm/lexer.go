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

	//lexer.tokens = make([][][]byte, len(lexer.sourceLines))

	// split all lines into tokens
	for _, line := range lexer.sourceLines {
		// filter comments delimited by '#'
		if isComment(line) || isSpaceLine(line) {
			continue
		}
		// split lines
		toks := bytes.FieldsFunc(line, func(c rune) bool {
			// ' ' || '\t' || ','
			return unicode.IsSpace(c) || c == ','
		})
		tokens := make([][]byte, len(toks))
		//lexer.tokens[lineIndex] = make([][]byte, len(toks))
		for tokIndex, tok := range toks {
			if value, ok := defines.FindRef(string(tok)); ok {
				tokens[tokIndex] = value
				continue
			}
			tokens[tokIndex] = tok
		}
		lexer.tokens = append(lexer.tokens, tokens)
	}
}

func (lexer *lexerCtx) Lines() [][]byte { return lexer.sourceLines }

func (lexer *lexerCtx) Tokens() [][][]byte { return lexer.tokens }

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

func isSpaceLine(line []byte) bool {
	return len(bytes.TrimSpace(line)) == 0
}

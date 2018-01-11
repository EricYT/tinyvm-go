package tdb

import (
	"bytes"
	"testing"
)

func TestTokenize(t *testing.T) {
	var foo []byte = []byte(" \t test      a b   \t c ")
	res := [][]byte{
		[]byte("test"),
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
	}
	tokens := tokenize(foo)
	if len(tokens) != len(res) {
		t.Fatalf("TestTokenize input foo(%#v) tokenize got wrong result(%#v)", foo, tokens)
	}
	for idx, token := range tokens {
		if !bytes.Equal(token, res[idx]) {
			t.Fatalf("TestTokenize foo %d token: %s not equal right one: %s", string(token), string(res[idx]))
		}
	}
}

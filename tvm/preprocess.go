package tvm

import (
	"bytes"
	"errors"
)

var (
	ErrorPreprocessDefineMissArgs error = errors.New("preprocess: define missing arguments")
)

func (c *Ctx) preprocess(source []byte) error {

	return nil
}

func processIncludes(source []byte) ([]byte, bool, error) {
	// find include delimiter
	directiveDelimiter := bytes.Index(source, []byte(TOK_INCLUDE))
	if directiveDelimiter == -1 {
		return source, false, nil
	}

	begin := directiveDelimiter
	end := bytes.IndexByte(source[begin:], '\n')

	//FIXME: there is a '\n' must be exists
	if end == -1 {
		return source, false, nil
	}

	// read include file
	filename := bytes.TrimSpace(source[begin+len(TOK_INCLUDE) : end])
	include, err := ReadFile(string(filename), ".vm")
	if err != nil {
		return source, false, err
	}

	// concate new source
	firstBlock := source[:begin]
	secondBlock := source[end+1:]

	// merge the include file
	var newsrc []byte
	newsrc = append(newsrc, firstBlock...)
	newsrc = append(newsrc, include...)
	newsrc = append(newsrc, secondBlock...)
	return newsrc, true, nil
}

type define struct {
	key   string
	value string
}

func processDefines(source []byte) ([]byte, *define, bool, error) {
	// find define delimiter
	directiveDelimiter := bytes.Index(source, []byte(TOK_DEFINE))
	if directiveDelimiter == -1 {
		return source, nil, false, nil
	}

	begin := directiveDelimiter
	end := bytes.IndexByte(source[begin:], '\n')

	//FIXME: there is a '\n' must be exists
	if end == -1 {
		return source, nil, false, nil
	}

	offset := len(TOK_DEFINE) + 1
	if begin+offset >= end {
		return source, nil, false, ErrorPreprocessDefineMissArgs
	}

	kv := bytes.SplitN(source[offset:end], []byte(" "), 2)
	key := string(bytes.TrimSpace(kv[0]))
	value := string(bytes.TrimSpace(kv[1]))

	// remove the line from source
	var remain []byte
	remain = append(remain, source[:begin]...)
	remain = append(remain, source[end+1:]...)
	return remain, &define{key, value}, true, nil
}

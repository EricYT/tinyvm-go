package tvm

import (
	"bytes"
	"errors"
	"log"
	"unicode"
)

var (
	ErrorPreprocessDefineMissArgs       error = errors.New("preprocess: define missing arguments")
	ErrorPreprocessDefineWrongArgs      error = errors.New("preprocess: define wrong arguments")
	ErrorPreprocessDefineMultipleDefine error = errors.New("preprocess: multiple define")
)

func (c *Ctx) preprocess(source []byte) ([]byte, error) {
	var included bool
	var def *define
	var err error

	for {
		source, included, err = processIncludes(source)
		if err != nil {
			return nil, err
		}
		source, def, err = processDefines(source)
		if err != nil {
			return nil, err
		}
		if def != nil {
			if _, ok := c.Prog.defines.FindRef(def.key); ok {
				log.Printf("Multiple definitions for %s.\n", def.key)
				return nil, ErrorPreprocessDefineMultipleDefine
			}
			c.Prog.defines.AddRef(def.key, def.value)
		}
		if !included && (def == nil) {
			break
		}
	}

	return source, nil
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
	end += begin

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
	value []byte
}

func processDefines(source []byte) ([]byte, *define, error) {
	// find define delimiter
	directiveDelimiter := bytes.Index(source, []byte(TOK_DEFINE))
	if directiveDelimiter == -1 {
		return source, nil, nil
	}

	begin := directiveDelimiter
	end := bytes.IndexByte(source[begin:], '\n')

	//FIXME: there is a '\n' must be exists
	if end == -1 {
		return source, nil, nil
	}
	end += begin

	offset := len(TOK_DEFINE)
	if begin+offset >= end {
		return source, nil, ErrorPreprocessDefineMissArgs
	}

	kv := bytes.FieldsFunc(source[begin+offset:end], func(c rune) bool {
		return unicode.IsSpace(c)
	})
	if len(kv) != 2 {
		return source, nil, ErrorPreprocessDefineWrongArgs
	}
	key := string(kv[0])
	value := kv[1]

	// remove the line from source
	var remain []byte
	remain = append(remain, source[:begin]...)
	remain = append(remain, source[end+1:]...)
	return remain, &define{key, value}, nil
}

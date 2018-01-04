package tvm

import (
	"io/ioutil"
	"log"
	"strings"
)

// Package tvm is the core package implement the tiny virtual machine.

/* core struct tvm */

type Ctx struct {
	prog *Prog
	mem  *Mem
}

func NewCtx() *Ctx {
	ctx := new(Ctx)

	return ctx
}

func (c *Ctx) Interpret(filename string) error {
	log.Printf("Prepare to interpret the file %s", filename)

	// Attempt to read the file all
	source, err := ReadFile(filename, ".vm")
	if err != nil {
		return err
	}

	// preprocess source
	if source, err = c.preprocess(source); err != nil {
		return err
	}

	// lexer analysis

	return nil
}

// utils
func ReadFile(filename string, extension string) ([]byte, error) {
	var source []byte
	var err error

	// read file by filename
	if source, err = ioutil.ReadFile(filename); err == nil {
		return source, nil
	}
	if strings.HasSuffix(filename, extension) {
		return nil, err
	}
	// try to read file with extension
	if source, err = ioutil.ReadFile(filename + extension); err != nil {
		return nil, err
	}
	return source, nil
}

package tvm

import (
	"bytes"
	"testing"
)

var jsrPreprocess []byte

func init() {
	jsrPreprocess, _ = ReadFile("./preprocessor/preprocess_data.vm", ".vm")
}

func TestProcessIncludes(t *testing.T) {
	src, err := ReadFile("./preprocessor/jsr.vm", ".vm")
	if err != nil {
		t.Fatalf("TestProcessIncludes jsr.vm read error: %s", err)
	}
	newSrc, included, err := processIncludes(src)
	if err != nil {
		t.Fatalf("TestProcessIncludes jsr.vm preprocessor error: %s", err)
	}
	if !included {
		t.Fatalf("TestProcessIncludes jsr.vm already include another one file")
	}
	if !bytes.Equal(newSrc, jsrPreprocess) {
		t.Fatalf("TestProcessIncludes jsr.vm preprocessor data not equal right one")
	}
}

func TestPrepreocessNotInclude(t *testing.T) {
	src, _ := ReadFile("./preprocessor/print_eax.vm", ".vm")
	newSrc, included, err := processIncludes(src)
	if err != nil {
		t.Fatalf("TestPrepreocessNotInclude preprocessor error: %s", err)
	}
	if included {
		t.Fatalf("TestPrepreocessNotInclude there is not include file exists.")
	}
	if !bytes.Equal(src, newSrc) {
		t.Fatalf("TestPrepreocessNotInclude there is no include file. preprocess shouldn't modify the original file.")
	}
}

func TestPreprocessDefines(t *testing.T) {
	src, _ := ReadFile("./preprocessor/define.vm", ".vm")
	newSrc, defines, defined, err := processDefines(src)
	if err != nil {
		t.Fatalf("TestPreprocessDefines process defines error: %s", err)
	}
	if defines == nil || !defined {
		t.Fatalf("TestPreprocessDefines not find defines")
	}
	if defines.key != "ONE" || defines.value != "1" {
		t.Fatalf("TestPreprocessDefines process define key '%%define' ONE 1 but got key/value %s %s", defines.key, defines.value)
	}
	indexFirstLine := bytes.IndexByte(src, '\n')
	if !bytes.Equal(newSrc, src[indexFirstLine+1:]) {
		t.Fatalf("TestPrepreocessNotInclude process defines not remove it from orignal source file")
	}
}

func TestPreprocessNoDefines(t *testing.T) {
	src, _ := ReadFile("./preprocessor/print_eax.vm", ".vm")
	_, _, defined, err := processDefines(src)
	if err != nil {
		t.Fatalf("TestPreprocessNoDefines process no define error: %s", err)
	}
	if defined {
		t.Fatalf("TestPreprocessNoDefines there is no defines in the file.")
	}
}

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
	newSrc, define, err := processDefines(src)
	if err != nil {
		t.Fatalf("TestPreprocessDefines process defines error: %s", err)
	}
	if define == nil {
		t.Fatalf("TestPreprocessDefines not find defines")
	}
	if define.key != "ONE" || string(define.value) != "1" {
		t.Fatalf("TestPreprocessDefines process define key '%%define ONE 1' but got key/value (%s) (%s)", define.key, string(define.value))
	}
	indexFirstLine := bytes.IndexByte(src, '\n')
	if !bytes.Equal(newSrc, src[indexFirstLine+1:]) {
		t.Fatalf("TestPreprocessDefines process defines not remove it from orignal source file")
	}

	// mutiple spaces between key and value
	src = []byte("   foo %define   foo    bar   \n")
	_, define, err = processDefines(src)
	if err != nil {
		t.Fatalf("TestPreprocessDefines process mutiple spaces between key and value data error: %s", err)
	}
	if define == nil {
		t.Fatalf("TestPreprocessDefines process not parse define.")
	}
	if define.key != "foo" || string(define.value) != "bar" {
		t.Fatalf("TestPreprocessDefines process define key '%%define   foo    bar' but got key/value (%s) (%s)", define.key, string(define.value))
	}

}

func TestPreprocessNoDefines(t *testing.T) {
	src, _ := ReadFile("./preprocessor/print_eax.vm", ".vm")
	_, define, err := processDefines(src)
	if err != nil {
		t.Fatalf("TestPreprocessNoDefines process no define error: %s", err)
	}
	if define != nil {
		t.Fatalf("TestPreprocessNoDefines there is no defines in the file.")
	}
}

func TestPreprocessDefinesWrong(t *testing.T) {
	src := "   %define\nbalabala"
	_, _, err := processDefines([]byte(src))
	if err != ErrorPreprocessDefineMissArgs {
		t.Fatalf("TestPreprocessDefinesWrong the define arguments is missing.")
	}

	src = "   %define \nbalabala"
	_, _, err = processDefines([]byte(src))
	if err != ErrorPreprocessDefineWrongArgs {
		t.Fatalf("TestPreprocessDefinesWrong the define arguments is missing key and value.")
	}

	src = "   %define x\nbalabala"
	_, _, err = processDefines([]byte(src))
	if err != ErrorPreprocessDefineWrongArgs {
		t.Fatalf("TestPreprocessDefinesWrong the define arguments is missing value.")
	}

	src = "foo\n%define x y z\nbalabala"
	_, _, err = processDefines([]byte(src))
	if err != ErrorPreprocessDefineWrongArgs {
		t.Fatalf("TestPreprocessDefinesWrong the define arguments is more than two key/value.")
	}
}

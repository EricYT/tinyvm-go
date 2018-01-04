package tvm

import (
	"bytes"
	"testing"
)

func TestHtabCtxAdd(t *testing.T) {
	htab := NewHtabCtx()

	if err := htab.Add("foo", 1); err != nil {
		t.Fatalf("htabCtx: add foo should not failed, error %s", err)
	}

	if err := htab.Add("foo", 1); err == nil {
		t.Fatalf("htabCtx: add foo again should failed.")
	}
}

func TestHtabCtxFind(t *testing.T) {
	htab := NewHtabCtx()

	htab.Add("foo", 1)

	if value, ok := htab.Find("foo"); !ok || value != 1 {
		t.Fatalf("htabCtx: foo not found or value should equal 1 but %d", value)
	}

	if _, ok := htab.Find("bar"); ok {
		t.Fatalf("htabNode: bar isn't in the htab but we got it")
	}
}

func TestHtabCtxAddRef(t *testing.T) {
	htab := NewHtabCtx()

	k := "foo"
	v := []byte("1")

	if err := htab.AddRef(k, v); err != nil {
		t.Fatalf("htabCtx: add ref foo should not failed, error %s", err)
	}

	if err := htab.AddRef(k, v); err == nil {
		t.Fatalf("htabCtx: add ref foo again should failed.")
	}
}

func TestHtabCtxFindRef(t *testing.T) {
	htab := NewHtabCtx()

	k := "foo"
	v := []byte("1")

	htab.AddRef(k, v)

	if value, ok := htab.FindRef(k); !ok || !bytes.Equal(value, v) {
		t.Fatalf("htabCtx: foo not found or value should be %v but %v", v, value)
	}

	if _, ok := htab.FindRef("bar"); ok {
		t.Fatalf("htabNode: the ref of bar isn't in the htab but we got it")
	}
}

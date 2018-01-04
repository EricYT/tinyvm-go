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

	if value, err := htab.Find("foo"); err != nil || value != 1 {
		t.Fatalf("htabCtx: find foo error %s value should be 1 but %d", err, value)
	}

	if _, err := htab.Find("bar"); err == nil {
		t.Fatalf("htabNode: find a not found key bar should error")
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

	if value, err := htab.FindRef(k); err != nil || !bytes.Equal(value, v) {
		t.Fatalf("htabCtx: find foo error %s value should be %v but %v", err, v, value)
	}

	if _, err := htab.FindRef("bar"); err == nil {
		t.Fatalf("htabNode: find a not found key bar should error")
	}
}

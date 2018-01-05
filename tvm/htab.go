package tvm

import "errors"

const (
	KEY_LENGTH int = 64
	HTAB_SIZE  int = 4096
)

var (
	ErrorHtabCtxKeyAlreadyExists error = errors.New("htab: key already exists")
	ErrorHtabCtxKeyNotFound      error = errors.New("htab: key not found")
)

type htabCtx struct {
	nodes map[string]*htabNode
}

func NewHtabCtx() *htabCtx {
	htab := new(htabCtx)

	htab.nodes = make(map[string]*htabNode)

	return htab
}

func (h *htabCtx) Add(key string, value int) error {
	if _, ok := h.nodes[key]; ok {
		return ErrorHtabCtxKeyAlreadyExists
	}
	h.nodes[key] = &htabNode{key: key, value: value}
	return nil
}

func (h *htabCtx) Find(key string) (*int, bool) {
	if node, ok := h.nodes[key]; ok {
		return &node.value, true
	}
	return nil, false
}

func (h *htabCtx) AddRef(key string, value []byte) error {
	if _, ok := h.nodes[key]; ok {
		return ErrorHtabCtxKeyAlreadyExists
	}
	h.nodes[key] = &htabNode{key: key, valueRaw: value}
	return nil
}

func (h *htabCtx) FindRef(key string) ([]byte, bool) {
	if node, ok := h.nodes[key]; ok {
		return node.valueRaw, true
	}
	return nil, false
}

type htabNode struct {
	key      string
	value    int
	valueRaw []byte
}

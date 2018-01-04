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

func (h *htabCtx) Find(key string) (int, error) {
	if node, ok := h.nodes[key]; ok {
		return node.value, nil
	}
	return -1, ErrorHtabCtxKeyNotFound
}

func (h *htabCtx) AddRef(key string, value []byte) error {
	if _, ok := h.nodes[key]; ok {
		return ErrorHtabCtxKeyAlreadyExists
	}
	h.nodes[key] = &htabNode{key: key, valueRaw: value}
	return nil
}

func (h *htabCtx) FindRef(key string) ([]byte, error) {
	if node, ok := h.nodes[key]; ok {
		return node.valueRaw, nil
	}
	return nil, ErrorHtabCtxKeyNotFound
}

type htabNode struct {
	key      string
	value    int
	valueRaw []byte
}

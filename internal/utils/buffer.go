package utils

import "sync"

type Buffer struct {
	Mux  sync.RWMutex
	Data []byte
	len  int
	cap  int
}

func (b *Buffer) Clear() {
	b.Mux.Lock()
	defer b.Mux.Unlock()
	b.len = 0
}

func (b *Buffer) Length() int {
	return b.len
}

func (b *Buffer) Append(data []byte) int {
	b.Mux.Lock()
	defer b.Mux.Unlock()
	if b.len+len(data) > b.cap {
		b.cap = b.len + len(data)
		b.Data = make([]byte, b.cap)
	}
	copy(b.Data[b.len:], data)
	b.len += len(data)
	return len(data)
}

func NewBuffer(cap int) *Buffer {
	return &Buffer{
		Data: make([]byte, cap),
		cap:  cap,
	}
}

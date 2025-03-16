package connectionmanager

import "sync"

type buffer struct {
	mux  sync.Mutex
	data []byte
	len  int
	cap  int
}

func (b *buffer) Clear() {
	b.mux.Lock()
	defer b.mux.Unlock()
	b.len = 0
}

func (b *buffer) Append(data []byte) int {
	b.mux.Lock()
	defer b.mux.Unlock()
	if b.len+len(data) > b.cap {
		b.cap = b.len + len(data)
		b.data = make([]byte, b.cap)
	}
	copy(b.data[b.len:], data)
	b.len += len(data)
	return len(data)
}

func NewBuffer(cap int) *buffer {
	return &buffer{
		data: make([]byte, cap),
		cap:  cap,
	}
}
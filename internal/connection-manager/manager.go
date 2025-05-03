package connectionmanager

import (
	"context"

	client "ipfs-store/internal/ipfs-client"

	"github.com/ipfs/go-cid"
)

const BUFFER_SIZE = 4 * 4096

type Record struct {
	Addr string
	Cid  cid.Cid
	Name string
}

type connection struct {
	ctx     context.Context
	bits    int
	health  bool
	cancel  context.CancelFunc
	manager *Manager
}

type Manager struct {
	client      *client.IPFSClient
	connections map[string]connection
	callbackch  chan Record
}

func NewManager(cbch chan Record) *Manager {
	return &Manager{
		client:      client.NewIPFSClient(),
		connections: make(map[string]connection),
		callbackch:  cbch,
	}
}

func (m *Manager) DialAddr(addr string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	connection := connection{ctx: ctx, cancel: cancel, manager: m}
	go func() {
		if err := Dial(string(addr), &connection, m.callbackch); err != nil {
			cancel()
		}
	}()
	m.connections[addr] = connection
	return nil
}

func (m *Manager) Close(addr string) {
	m.connections[addr].cancel()
}

func (m *Manager) Health(addr string) bool {
	return m.connections[addr].health
}

func (m *Manager) GetBits(addr string) int {
	return m.connections[addr].bits
}

func (c *connection) SetHealthToFalse() {
	c.health = false
}

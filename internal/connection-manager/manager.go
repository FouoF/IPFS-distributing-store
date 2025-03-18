package connectionmanager

import (
	"context"

	"github.com/ipfs/go-cid"
	client "ipfs-store/internal/ipfs-client"
)

const BUFFER_SIZE = 4096

type Record struct {
	Addr    address
	Cid     cid.Cid
	Name	string
}

type address string

type connection struct {
	ctx 	context.Context
	bits    int
	health 	bool
	cancel  context.CancelFunc
	manager *Manager
}

type Manager struct {
	client 	   *client.IPFSClient
	connections map[address]connection
	callbackch  chan Record
}

func NewManager(cbch chan Record) *Manager {
	return &Manager{
		client: 	 client.NewIPFSClient(),
		connections: make(map[address]connection),
		callbackch:  cbch,
	}
}

func (m *Manager) DialAddr(addr address) error{
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	connection := connection{ctx: ctx, cancel: cancel, manager: m}
	if err := Dial(string(addr), &connection, m.callbackch); err != nil {
		cancel()
		return err
	}
	m.connections[addr] = connection
	return nil
}

func (m *Manager) Close(addr address){
	m.connections[addr].cancel()
}

func (m *Manager) Health(addr address) bool{
	return m.connections[addr].health
}

func (m *Manager) GetBits(addr address) int{
	return m.connections[addr].bits
}
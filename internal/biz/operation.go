package biz

import (
	"context"
	"errors"
	"sync"

	v1 "ipfs-store/api/admin-service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type EndpointList struct {
	endpoints map[string]*v1.Endpoint
	lock      sync.RWMutex
}

type NodeList struct {
	nodes map[string]*v1.Node
	lock  sync.RWMutex
}

type OperationUsecase struct {
	endpointList EndpointList
	nodeList     NodeList
	log          *log.Helper
}

func NewOperationUsecase(logger log.Logger) *OperationUsecase {
	return &OperationUsecase{
		log: log.NewHelper(logger), 
		endpointList: EndpointList{endpoints: make(map[string]*v1.Endpoint)}, 
		nodeList: NodeList{nodes: make(map[string]*v1.Node)}}
}

func (uc *OperationUsecase) AddEndpoint(ctx context.Context, endpoint *v1.Endpoint) error {
	uc.log.WithContext(ctx).Infof("add endpoint: %v", endpoint)
	uc.endpointList.lock.Lock()
	if uc.endpointList.endpoints[endpoint.Addr] != nil {
		return errors.New("endpoint already exists")
	}
	defer uc.endpointList.lock.Unlock()
	uc.endpointList.endpoints[endpoint.Addr] = &v1.Endpoint{
		Addr:        endpoint.Addr,
		Description: endpoint.Description,
		Id:          endpoint.Id,
		Index:       endpoint.Index,
	}
	return nil
}

func (uc *OperationUsecase) ListEndpoints(ctx context.Context) ([]*v1.Endpoint, error) {
	uc.log.WithContext(ctx).Infof("list endpoints")
	uc.endpointList.lock.RLock()
	defer uc.endpointList.lock.RUnlock()
	var eps []*v1.Endpoint
	for _, endpoint := range uc.endpointList.endpoints {
		eps = append(eps, endpoint)
	}
	return eps, nil
}

func (uc *OperationUsecase) GetEndpoint(ctx context.Context, addr string) (*v1.Endpoint, error) {
	uc.log.WithContext(ctx).Infof("get endpoint: %v", addr)
	uc.endpointList.lock.RLock()
	defer uc.endpointList.lock.RUnlock()
	return uc.endpointList.endpoints[addr], nil
}

func (uc *OperationUsecase) DeleteEndpoint(ctx context.Context, addr string) error {
	uc.log.WithContext(ctx).Infof("delete endpoint: %v", addr)
	uc.endpointList.lock.Lock()
	defer uc.endpointList.lock.Unlock()
	delete(uc.endpointList.endpoints, addr)
	return nil
}
func (uc *OperationUsecase) AddNode(ctx context.Context, node *v1.Node) error {
	uc.log.WithContext(ctx).Infof("add node: %v", node)
	uc.nodeList.lock.Lock()
	defer uc.nodeList.lock.Unlock()
	if uc.nodeList.nodes[node.Addr] != nil {
		log.Errorf("node already exists")
		return nil
	}
	uc.nodeList.nodes[node.Addr] = &v1.Node{
		Addr: node.Addr,
		Id:   node.Id,
		Name: node.Name,
	}
	return nil
}

func (uc *OperationUsecase) ListNodes(ctx context.Context) ([]*v1.Node, error) {
	uc.log.WithContext(ctx).Infof("list Nodes")
	uc.nodeList.lock.RLock()
	defer uc.nodeList.lock.RUnlock()
	var nodes []*v1.Node
	for _, node := range uc.nodeList.nodes {
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func (uc *OperationUsecase) GetNode(ctx context.Context, addr string) (*v1.Node, error) {
	uc.log.WithContext(ctx).Infof("Get Node: %v", addr)
	uc.nodeList.lock.RLock()
	defer uc.nodeList.lock.RUnlock()
	return uc.nodeList.nodes[addr], nil
}

func (uc *OperationUsecase) DeleteNode(ctx context.Context, addr string) error {
	uc.log.WithContext(ctx).Infof("delete node: %v", addr)
	uc.nodeList.lock.Lock()
	defer uc.nodeList.lock.Unlock()
	delete(uc.nodeList.nodes, addr)
	return nil
}
package service

import (
	"context"
	v1 "ipfs-store/api/admin-service/v1"
	"ipfs-store/internal/biz"
	manager "ipfs-store/internal/connection-manager"
)

type OperationService struct {
	v1.UnimplementedOperationServer

	cm *manager.Manager
	uc *biz.OperationUsecase
	du *biz.DataUsecase
	ch chan manager.Record
}

func NewOperationService(uc *biz.OperationUsecase, du *biz.DataUsecase) *OperationService {
	ch := make(chan manager.Record)
	manager := manager.NewManager(ch)
	s := &OperationService{uc: uc, du: du, ch: ch, cm: manager}
	go WatchCh(s)
	s.AddEndpoint(context.TODO(), &v1.AddEndpointRequest{Endpoint: 
		&v1.Endpoint{Addr: "endpoint.default.svc.cluster.local:50051", 
		Index: &v1.Index{Name: "心率", L1: "1号房间", L2: "1号床", Leafname: ""}} })
	return s
}

func (s *OperationService) ListNode(ctx context.Context, req *v1.ListNodeRequest) (*v1.ListNodeReply, error) {
	nodes, err := s.uc.ListNodes(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListNodeReply{Node: nodes}, nil
}

func (s *OperationService) AddNode(ctx context.Context, req *v1.AddNodeRequest) (*v1.AddNodeReply, error) {
	n := v1.Node{Id: req.Node.Id, Name: req.Node.Name, Addr: req.Node.Addr, Status: v1.Status_OFFLINE}
	err := s.uc.AddNode(ctx, &n)
	if err != nil {
		return nil, err
	}
	return &v1.AddNodeReply{}, nil
}

func (s *OperationService) RemoveNode(ctx context.Context, req *v1.RemoveNodeRequest) (*v1.RemoveNodeReply, error) {
	err := s.uc.DeleteNode(ctx, req.Addr)
	if err != nil {
		return nil, err
	}
	return &v1.RemoveNodeReply{}, nil
}

func (s *OperationService) ListEndpoint(ctx context.Context, req *v1.ListEndpointRequest) (*v1.ListEndpointReply, error) {
	endpoints, err := s.uc.ListEndpoints(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListEndpointReply{Endpoint: endpoints}, nil
}

func (s *OperationService) GetEndpoint(ctx context.Context, req *v1.GetEndpointRequest) (*v1.GetEndpointReply, error) {
	node, err := s.du.GetNode(ctx, s.du.V1ToDatastore(req.Index))
	if err != nil {
		return nil, err

	}
	endpoint, err := s.uc.GetEndpoint(ctx, node.Name)
	var endpoints []*v1.Endpoint
	endpoints = append(endpoints, endpoint)
	if err != nil {
		return nil, err
	}
	return &v1.GetEndpointReply{Endpoint: endpoints}, nil
}

func (s *OperationService) AddEndpoint(ctx context.Context, req *v1.AddEndpointRequest) (*v1.AddEndpointReply, error) {
	err := s.uc.AddEndpoint(ctx, req.Endpoint)
	if err != nil {
		return nil, err
	}
	err = s.cm.DialAddr(req.Endpoint.Addr)
	if err != nil {
		return nil, err
	}
	return &v1.AddEndpointReply{}, nil
}

func (s *OperationService) RemoveEndpoint(ctx context.Context, req *v1.RemoveEndpointRequest) (*v1.RemoveEndpointReply, error) {
	err := s.uc.DeleteEndpoint(ctx, req.Addr)
	if err != nil {
		return nil, err
	}
	s.cm.Close(req.Addr)
	return &v1.RemoveEndpointReply{}, nil
}

func (s *OperationService) CreateIndex(ctx context.Context, req *v1.CreateIndexRequest) (*v1.CreateIndexReply, error) {
	didx := s.du.V1ToDatastore(req.Index)
	err := s.du.AddIndex(ctx, didx)
	if err != nil {
		return nil, err
	}
	return &v1.CreateIndexReply{}, nil
}

func (s *OperationService) ListIndex(ctx context.Context, req *v1.ListIndexRequest) (*v1.ListIndexReply, error) {
	didx := s.du.V1ToDatastore(req.Index)
	node, err := s.du.GetNode(ctx, didx)
	indices := make([]string, 0)
	for _, v := range node.Children {
		indices = append(indices, v.Name)
	}
	if err != nil {
		return nil, err
	}
	return &v1.ListIndexReply{Index: indices}, nil
}

func WatchCh(s *OperationService) {
	ctx := context.TODO()
	for {
		record := <-s.ch
		endpoint, _ := s.uc.GetEndpoint(ctx, string(record.Addr))
		index := s.du.V1ToDatastore(endpoint.Index)
		index.LeafName = record.Name
		s.du.AddLeaf(ctx, index, record.Cid)
	}
}
package service

import (
	"context"
	v1 "ipfs-store/api/admin-service/v1"
	"ipfs-store/internal/biz"
)

type OperationService struct {
	v1.UnimplementedOperationServer

	uc *biz.OperationUsecase
	du *biz.DataUsecase
}

func NewOperationService(uc *biz.OperationUsecase, du *biz.DataUsecase) *OperationService {
	return &OperationService{uc: uc, du: du}
}

func (s *OperationService) ListNode(ctx context.Context, req *v1.ListNodeRequest) (*v1.ListNodeReply, error) {
	nodes, err := s.uc.ListNodes(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListNodeReply{Node: nodes}, nil
}

func (s *OperationService) AddNode(ctx context.Context, req *v1.AddNodeRequest) (*v1.AddNodeReply, error) {
	n := v1.Node{Name: req.Node.Name, Addr: req.Node.Addr, Status: v1.Status_OFFLINE}
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
	return &v1.AddEndpointReply{}, nil
}

func (s *OperationService) RemoveEndpoint(ctx context.Context, req *v1.RemoveEndpointRequest) (*v1.RemoveEndpointReply, error) {
	didx := s.du.V1ToDatastore(req.Index)
	node, err := s.du.GetNode(ctx ,didx)
	if err != nil {
		return nil, err
	}
	err = s.uc.DeleteEndpoint(ctx, node.Name)
	if err != nil {
		return nil, err
	}
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

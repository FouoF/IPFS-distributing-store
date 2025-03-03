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
	// 实现逻辑：列出节点
	nodes, err := s.uc.ListNodes(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListNodeReply{Node: nodes}, nil
}

func (s *OperationService) AddNode(ctx context.Context, req *v1.AddNodeRequest) (*v1.AddNodeReply, error) {
	n := v1.Node{Name: req.Name, Addr: req.Addr, Status: v1.Status_OFFLINE}
	err := s.uc.AddNode(ctx, &n)
	if err != nil {
		return nil, err
	}
	return &v1.AddNodeReply{Success: true}, nil
}

func (s *OperationService) RemoveNode(ctx context.Context, req *v1.RemoveNodeRequest) (*v1.RemoveNodeReply, error) {
	err := s.uc.DeleteNode(ctx, req.Addr)
	if err != nil {
		return nil, err
	}
	return &v1.RemoveNodeReply{Success: true}, nil
}

func (s *OperationService) ListEndpoint(ctx context.Context, req *v1.ListEndpointRequest) (*v1.ListEndpointReply, error) {
	endpoints, err := s.uc.ListEndpoints(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListEndpointReply{Endpoint: endpoints}, nil
}

func (s *OperationService) GetEndpoint(ctx context.Context, req *v1.GetEndpointRequest) (*v1.GetEndpointReply, error) {
	node, err := s.du.GetNode(ctx, s.du.V1ToDatastore(req))
	endpoint, err := s.uc.GetEndpoint(ctx, node.Name)
	var endpoints []*v1.Endpoint
	endpoints = append(endpoints, endpoint)
	if err != nil {
		return nil, err
	}
	return &v1.GetEndpointReply{Endpoint: endpoints}, nil
}

func (s *OperationService) AddEndpoint(ctx context.Context, req *v1.AddEndpointRequest) (*v1.AddEndpointReply, error) {
	endpoint := &v1.Endpoint{Name: req.Name, Addr: req.Addr, Desctiption: req.Desctiption}
	err := s.uc.AddEndpoint(ctx, endpoint)
	if err != nil {
		return nil, err
	}
	return &AddEndpointReply{EndpointId: endpointID}, nil
}

func (s *OperationService) RemoveEndpoint(ctx context.Context, req *RemoveEndpointRequest) (*RemoveEndpointReply, error) {
	// 实现逻辑：删除端点
	err := s.du.RemoveEndpoint(ctx, req.EndpointId)
	if err != nil {
		return nil, err
	}
	return &RemoveEndpointReply{}, nil
}

func (s *OperationService) CreateIndex(ctx context.Context, req *CreateIndexRequest) (*CreateIndexReply, error) {
	// 实现逻辑：创建索引
	indexID, err := s.du.CreateIndex(ctx, req.Index)
	if err != nil {
		return nil, err
	}
	return &CreateIndexReply{IndexId: indexID}, nil
}

func (s *OperationService) ListIndex(ctx context.Context, req *ListIndexRequest) (*ListIndexReply, error) {
	// 实现逻辑：列出索引
	indices, err := s.du.ListIndices(ctx)
	if err != nil {
		return nil, err
	}
	return &ListIndexReply{Indices: indices}, nil
}

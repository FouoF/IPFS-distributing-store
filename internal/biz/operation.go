package biz

import (
	"context"

	v1 "ipfs-store/api/admin-service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type OperationUsecase struct {
	endpoints map[string]v1.Endpoint
	nodes     map[string]v1.Node
	log       *log.Helper
}

func NewOperationUsecase(logger log.Logger) *OperationUsecase {
	return &OperationUsecase{log: log.NewHelper(logger)}
}

func (uc *OperationUsecase) AddEndpoint(ctx context.Context, endpoint *v1.Endpoint) error {
	uc.log.WithContext(ctx).Infof("add endpoint: %v", endpoint)
	uc.endpoints[endpoint.Addr] = v1.Endpoint{
		Addr:        endpoint.Addr,
		Desctiption: endpoint.Desctiption,
		Id:          endpoint.Id,
		L1Idx:       endpoint.L1Idx,
		L2Idx:       endpoint.L2Idx,
		L3Idx:       endpoint.L3Idx,
	}
	return nil
}
func (uc *OperationUsecase) AddNode(ctx context.Context, node *v1.Node) error {
	uc.log.WithContext(ctx).Infof("add node: %v", node)
	uc.nodes[node.Addr] = v1.Node{
		Addr: node.Addr,
		Id:   node.Id,
		Name: node.Name,
	}
	return nil
}
func (uc *OperationUsecase) DialEndpoint(ctx context.Context, addr string) error {
	uc.log.WithContext(ctx).Infof("dial endpoint: %v", addr)
	return nil
}

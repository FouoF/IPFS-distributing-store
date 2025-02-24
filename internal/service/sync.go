package service

import (
	"context"

	v1 "ipfs-store/api/admin-service/v1"
	"ipfs-store/internal/biz"
)

// GreeterService is a greeter service.
type SyncService struct {
	v1.UnimplementedSyncServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewSyncService(uc *biz.GreeterUsecase) *OperationService {
	return &OperationService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

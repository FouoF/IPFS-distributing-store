package service

import (
	v1 "ipfs-store/api/admin-service/v1"
	"ipfs-store/internal/biz"
)

type OperationService struct {
	v1.UnimplementedOperationServer

	uc *biz.OperationUsecase
	du *biz.DataUsecase
	su *biz.SyncUsecase
}

func NewOperationService(uc *biz.OperationUsecase, du *biz.DataUsecase, su *biz.SyncUsecase) *OperationService {
	return &OperationService{uc: uc, du: du, su: su}
}

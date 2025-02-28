package biz

import (
	v1 "ipfs-store/api/admin-service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type SyncUsecase struct {
	endpoints []v1.Endpoint
	nodes     []v1.Node
	log       *log.Helper
}

func NewSyncUsecase(logger log.Logger) *SyncUsecase {
	return &SyncUsecase{log: log.NewHelper(logger)}
}

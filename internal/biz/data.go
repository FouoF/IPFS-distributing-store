package biz

import (
	"context"

	v1 "ipfs-store/api/admin-service/v1"
	"ipfs-store/internal/datastore"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_RESOURCE_NOT_FOUND.String(), "resource not found")
)

type DataRepo interface {
	AddIndex(ctx context.Context, idx datastore.Index) error
	GetNode(ctx context.Context, idx datastore.Index) *datastore.Node
	AddLeaf(ctx context.Context, idx datastore.Index, cid string) error
}

type DataUsecase struct {
	repo DataRepo
	log  *log.Helper
}

func NewDataUsecase(repo DataRepo, logger log.Logger) *DataUsecase {
	return &DataUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DataUsecase) AddIndex(ctx context.Context, idx datastore.Index) error {
	uc.log.WithContext(ctx).Infof("addindex: %v", idx)
	return uc.repo.AddIndex(ctx, idx)
}

func (uc *DataUsecase) GetNode(ctx context.Context, idx datastore.Index) (*datastore.Node, error) {
	uc.log.WithContext(ctx).Infof("getnode: %v", idx)
	var node *datastore.Node
	if node = uc.repo.GetNode(ctx, idx); node == nil {
		return nil, ErrUserNotFound
	}
	return node, nil
}

func (uc *DataUsecase) AddLeaf(ctx context.Context, idx datastore.Index, cid string) error {
	uc.log.WithContext(ctx).Infof("addleaf: %v", idx)
	return uc.repo.AddLeaf(ctx, idx, cid)
}

func (uc *DataUsecase) V1ToDatastore(idx *v1.Index) datastore.Index {
	return datastore.Index{Name: idx.Name, L1: idx.L1, L2: idx.L2}
}

func (uc *DataUsecase) DatastoreToV1(idx datastore.Index) *v1.Index {
	return &v1.Index{Name: idx.Name, L1: idx.L1, L2: idx.L2}
}
package data

import (
	"context"

	"ipfs-store/internal/biz"
	"ipfs-store/internal/datastore"

	"github.com/go-kratos/kratos/v2/log"
)

type DataRepo struct {
	data *datastore.TreeStore
	log  *log.Helper
}

func NewDataRepo(data *Data, logger log.Logger) biz.DataRepo {
	return &DataRepo{
		data: data.Datastore,
		log:  log.NewHelper(logger),
	}
}

func (r *DataRepo) AddIndex(ctx context.Context, idx datastore.Index) error {
	return r.data.AddIndex(idx)
}

func (r *DataRepo) GetNode(ctx context.Context, idx datastore.Index) *datastore.Node {
	return r.data.GetNode(idx)
}

func (r *DataRepo) AddLeaf(ctx context.Context, idx datastore.Index, cid string) error {
	return r.data.AddLeaf(idx, cid)
}

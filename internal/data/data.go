package data

import (
	"ipfs-store/internal/conf"
	"ipfs-store/internal/datastore"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDataRepo)

// Data .
type Data struct {
	Datastore *datastore.TreeStore
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{Datastore: datastore.NewTreeStore()}, cleanup, nil
}

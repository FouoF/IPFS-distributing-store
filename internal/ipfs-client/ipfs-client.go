package ipfsclient

import (
	"github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/client/rpc"
)

type IPFSClient struct {
	rpcClient rpc.HttpApi
}

func Pinning(data []byte) (cid.Cid, error) {
	return cid.Cid{}, nil
}
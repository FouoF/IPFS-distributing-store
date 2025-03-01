package ipfsclient

import (
	"github.com/ipfs/kubo/client/rpc"
)

type IPFSClient struct {
	rpcClient rpc.HttpApi
}
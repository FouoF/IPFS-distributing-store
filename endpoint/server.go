// server.go
package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	v1 "ipfs-store/api/admin-service/v1"
)

type server struct {
	v1.UnimplementedSyncServer
}

func (s *server) SyncDataFromEndpoint(req *v1.SyncDataFromEndpointRequest, stream v1.Sync_SyncDataFromEndpointServer) error {
	for i := 0; i < 5; i++ {
		fileChunk := &v1.FileChunk{
			L1Idx:        "level1",
			L2Idx:        "level2",
			L3Idx:        "level3",
			ChunkIndex:   int64(i),
			Data:         []byte(fmt.Sprintf("file chunk %d", i)),
			IsFinalChunk: i == 4,
		}
		if err := stream.Send(fileChunk); err != nil {
			return err
		}
	}
	return nil
}
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	v1.RegisterSyncServer(s, &server{})

	fmt.Println("Server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

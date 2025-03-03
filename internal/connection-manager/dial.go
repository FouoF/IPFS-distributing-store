package connectionmanager

import (
	"context"
	"fmt"
	"io"
	v1 "ipfs-store/api/admin-service/v1"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Dial(target string, ctx context.Context) {
	// 连接到 gRPC 服务器
	conn, err := grpc.NewClient(fmt.Sprintf("dns:///%v", target), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := v1.NewSyncClient(conn)

	// 客户端调用 SyncDataFromEndpoint 方法
	stream, err := client.SyncDataFromEndpoint(ctx, &v1.SyncDataFromEndpointRequest{})
	if err != nil {
		log.Fatalf("could not sync data: %v", err)
	}

	for {
		fileChunk, err := stream.Recv()
		if err == io.EOF {
			break // 结束接收
		}
		if err != nil {
			log.Fatalf("could not receive file chunk: %v", err)
		}
		fmt.Printf("Received chunk: %s\n", string(fileChunk.Data))
	}
}

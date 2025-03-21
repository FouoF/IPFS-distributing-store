// server.go
package main

import (
	"fmt"
	v1 "ipfs-store/api/admin-service/v1"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	v1.UnimplementedSyncServer
}

func (s *server) SyncDataFromEndpoint(req *v1.SyncDataFromEndpointRequest, stream v1.Sync_SyncDataFromEndpointServer) error {
	// 示例索引
	idx := v1.Index{
		Name:     "心率",
		L1:       "1号房间",
		L2:       "1号床",
		Leafname: "2025.3.19 15:00",
	}

	// 假设数据从 2025年3月19日 15:00 开始
	startTime := time.Date(2025, 3, 19, 15, 0, 0, 0, time.UTC)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for j := 0; j < 5; j++ {
		// 模拟心率数据：每个数据块代表 1 分钟间隔的心率数据
		for i := range 60 {
			// 每个数据块的心率数据为：时间戳 + 心率值
			startTime = startTime.Add(time.Minute * time.Duration(1))
			heartRate := 60 + r.Intn(10) // 假设心率数据从 60 开始，并且每次递增 1

			// 创建一个 FileChunk
			fileChunk := &v1.FileChunk{
				Index:        &idx,
				ChunkIndex:   int64(i),
				Data:         fmt.Appendf(nil, "%s 心率：%d bpm", startTime.Format("2006-01-02 15:04:05"), heartRate),
				IsFinalChunk: i == 59, // 最后一块数据设置 IsFinalChunk 为 true
			}

			// 发送数据块
			if err := stream.Send(fileChunk); err != nil {
				return err
			}
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

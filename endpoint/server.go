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

func (s *server) dataTemplate(name, l1, l2, unit string, base int, stream v1.Sync_SyncDataFromEndpointServer) error {
	startTime := time.Date(2025, 3, 19, 15, 0, 0, 0, time.UTC)
	idx := v1.Index{
		Name:     name,
		L1:       l1,
		L2:       l2,
		Leafname: startTime.Format("2006-01-02 15:04:05"),
	}

	// 假设数据从 2025年3月19日 15:00 开始
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for j := 0; j < 5; j++ {
		idx.Leafname = startTime.Format("2006-01-02 15:04:05")
		// 模拟心率数据：每个数据块代表 1 分钟间隔的心率数据
		for i := range 60 {
			// 每个数据块的心率数据为：时间戳 + 心率值
			startTime = startTime.Add(time.Minute * time.Duration(1))
			value := base + r.Intn(10)

			// 创建一个 FileChunk
			fileChunk := &v1.FileChunk{
				Index:        &idx,
				ChunkIndex:   int64(i),
				Data:         fmt.Appendf(nil, "%s %s:%d %s\n", startTime.Format("2006-01-02 15:04:05"), name, value, unit),
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

func (s *server) SyncDataFromEndpoint(req *v1.SyncDataFromEndpointRequest, stream v1.Sync_SyncDataFromEndpointServer) error {
	var err error
	err = s.dataTemplate("心率", "1号房间", "1号床", "次/分", 60, stream)
	if err != nil {
		return err
	}
	err = s.dataTemplate("心率", "2号房间", "1号床", "次/分", 60, stream)
	if err != nil {
		return err
	}
	err = s.dataTemplate("心率", "1号房间", "2号床", "次/分", 60, stream)
	if err != nil {
		return err
	}
	err = s.dataTemplate("心率", "2号房间", "2号床", "次/分", 60, stream)
	if err != nil {
		return err
	}
	err = s.dataTemplate("血压", "1号房间", "1号床", "/120mmHg", 70, stream)
	if err != nil {
		return err
	}
	err = s.dataTemplate("血压", "2号房间", "1号床", "/120mmHg", 70, stream)
	if err != nil {
		return err
	}
	err = s.dataTemplate("血压", "1号房间", "2号床", "/120mmHg", 70, stream)
	if err != nil {
		return err
	}
	err = s.dataTemplate("血压", "2号房间", "2号床", "/120mmHg", 70, stream)
	if err != nil {
		return err
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

// server.go
package main

import (
	"fmt"
	"io"
	v1 "ipfs-store/api/admin-service/v1"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	v1.UnimplementedSyncServer
}

func splitVideoFile(videoPath string, chunkSize int64) ([][]byte, error) {
	var chunks [][]byte

	// 打开视频文件
	file, err := os.Open(videoPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open video file: %v", err)
	}
	defer file.Close()

	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}
	fileSize := fileInfo.Size()

	// 逐块读取文件并分段
	buf := make([]byte, chunkSize)
	for offset := int64(0); offset < fileSize; offset += chunkSize {
		n, err := file.ReadAt(buf, offset)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("failed to read file chunk: %v", err)
		}
		chunks = append(chunks, buf[:n])
	}

	return chunks, nil
}

func (s *server) dataTemplate(name, l1, l2, path string, stream v1.Sync_SyncDataFromEndpointServer) error {
	startTime := time.Date(2025, 3, 19, 15, 0, 0, 0, time.UTC)
	idx := v1.Index{
		Name:     name,
		L1:       l1,
		L2:       l2,
		Leafname: startTime.Format("2006-01-02 15:04:05"),
	}

	// 假设视频路径为指定路径（这里使用一个示例路径，实际路径根据需求指定）
	videoPath := path  + "video.mp4"
	chunkSize := int64(4096) // 每个数据块大小为 1MB（可以根据需求调整）

	// 获取视频文件分段
	chunks, err := splitVideoFile(videoPath, chunkSize)
	if err != nil {
		return fmt.Errorf("failed to split video file: %v", err)
	}

	// 模拟数据发送，分块发送视频数据
	for j := 0; j < 5; j++ {
		idx.Leafname = startTime.Format("2006-01-02 15:04:05")

		for i, chunk := range chunks {
			startTime = startTime.Add(time.Minute * time.Duration(1))

			// 创建 FileChunk
			fileChunk := &v1.FileChunk{
				Index:        &idx,
				ChunkIndex:   int64(i),
				Data:         chunk,  // 这里直接传递视频块数据
				IsFinalChunk: i == len(chunks)-1, // 最后一块数据设置 IsFinalChunk 为 true
			}

			// 发送数据块
			if err := stream.Send(fileChunk); err != nil {
				return err
			}

			fmt.Println(idx.Leafname)
		}
	}

	time.Sleep(1 * time.Second)
	return nil
}


func (s *server) SyncDataFromEndpoint(req *v1.SyncDataFromEndpointRequest, stream v1.Sync_SyncDataFromEndpointServer) error {
	err := s.dataTemplate("监控", "一号楼", "1号房间", "/data", stream)
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

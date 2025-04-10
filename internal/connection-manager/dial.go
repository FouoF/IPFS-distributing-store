package connectionmanager

import (
	"fmt"
	"io"
	v1 "ipfs-store/api/admin-service/v1"
	"ipfs-store/internal/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Dial(target string, connection *connection, ch chan Record) error {
	conn, err := grpc.NewClient(fmt.Sprintf("dns:///%v", target), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := v1.NewSyncClient(conn)

	connection.health = true
	defer connection.SetHealthToFalse()
	stream, err := client.SyncDataFromEndpoint(connection.ctx, &v1.SyncDataFromEndpointRequest{})
	if err != nil {
		return err
	}

	buffer := utils.NewBuffer(BUFFER_SIZE)

	for {
		fileChunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		connection.bits = connection.bits + buffer.Append(fileChunk.Data)
		if fileChunk.IsFinalChunk {
			name := fileChunk.Index.Leafname
			cid, err := connection.manager.client.PinDirect(buffer, name)
			if err != nil {
				fmt.Printf("pin direct error: %v", err)
				continue
			}
			ch <- Record{Cid: cid, Addr: target, Name: name}
			buffer.Clear()
		}
	}

	return nil
}

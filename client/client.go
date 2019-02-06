package client

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/faroyam/rpcftp/rpcftp"
	"github.com/faroyam/rpcftp/utils"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Client ...
type Client struct {
	remoteAddr string
	c          pb.RpcftpClient
}

// NewClient returns Client instance with server addr specified
func NewClient(addr, port string) (*Client, error) {
	fullAddr := addr + ":" + port
	conn, err := grpc.Dial(fullAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewRpcftpClient(conn)

	return &Client{remoteAddr: fullAddr, c: c}, nil
}

// UploadFile uloads file to the server using Upload rpc
func (cl *Client) UploadFile(localFileName, remoteFileName string) error {
	ctx := context.Background()
	file, err := os.Open(localFileName)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer file.Close()

	stream, err := cl.c.Upload(ctx)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	startTime := time.Now()
	buf := make([]byte, 1024)
	load := utils.LoadProgress()
	for {
		fmt.Print(load())
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err.Error())
			return err
		}
		if err := stream.Send(&pb.Chunk{
			FileName: remoteFileName,
			Content:  buf[:n],
		}); err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err.Error())
			return err
		}
	}
	elapsedTime := time.Since(startTime)
	uploadStatus, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("\nElapsed time: %s\nServer answer: %s\nCode: %s\n", elapsedTime, uploadStatus.Message, uploadStatus.Code)
	return err
}

// DownloadFile dowloads file from the server using Download rpc
func (cl *Client) DownloadFile(remoteFileName, localFileName string) error {

	if _, err := os.Stat(localFileName); !os.IsNotExist(err) {
		fmt.Println(localFileName + " is already exists")
		return err
	}

	err := utils.CreateDirs(localFileName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	file, err := os.Create(localFileName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	ctx := context.Background()

	stream, err := cl.c.Download(ctx, &pb.DownloadRequest{Path: remoteFileName})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	startTime := time.Now()
	load := utils.LoadProgress()
	for {
		fmt.Print(load())
		chunk, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err.Error())
			os.Remove(localFileName)
			return err
		}
		if _, err := file.Write(chunk.Content); err != nil {
			fmt.Println(err.Error())
			os.Remove(localFileName)
			return err
		}
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("\nElapsed time: %s\nDone.\n", elapsedTime)
	return nil
}

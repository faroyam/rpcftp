package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	pb "github.com/faroyam/rpcftp/rpcftp"
	"github.com/faroyam/rpcftp/utils"

	"google.golang.org/grpc/status"
)

// Server ...
type Server struct {
	port    string
	rootDir string
}

// NewServer returns Server instance with port and serve directory specified
func NewServer(port, rootDir string) *Server {
	return &Server{port: port, rootDir: rootDir}
}

// Upload implement Upload rpc
func (s *Server) Upload(stream pb.Rpcftp_UploadServer) error {
	chunk, err := stream.Recv()
	log.Printf("upload rpc called: %v%v", s.rootDir, chunk.FileName)
	if err != nil {
		log.Println(err.Error())
		st := status.New(codes.Canceled, "connection error")
		return st.Err()
	}

	localFileName := path.Join(s.rootDir, chunk.FileName)
	if _, err := os.Stat(localFileName); !os.IsNotExist(err) {
		log.Println(localFileName + " is already exists")
		stream.SendAndClose(&pb.UploadStatus{
			Message: localFileName + " is already exists",
			Code:    pb.UploadStatusCode_Failed,
		})
		return err
	}

	err = utils.CreateDirs(localFileName)
	if err != nil {
		fmt.Println(err.Error())
		uploadErrHandle("creating dirs error", stream, err)
		return err
	}

	file, err := os.Create(localFileName)
	if err != nil {
		fmt.Println(err)
		uploadErrHandle("creating file error", stream, err)
		return err
	}
	defer file.Close()

	if _, err := file.Write(chunk.Content); err != nil {
		log.Println(err.Error())
		uploadErrHandle("writing file error", stream, err)
		return err
	}

	for {
		chunk, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			uploadErrHandle("reading from stream error", stream, err)
			return err
		}
		if _, err := file.Write(chunk.Content); err != nil {
			uploadErrHandle("writing file error", stream, err)
			return err
		}
	}
	stream.SendAndClose(&pb.UploadStatus{
		Message: "success",
		Code:    pb.UploadStatusCode_Ok,
	})
	return nil
}

// Download implement Download rpc
func (s *Server) Download(in *pb.DownloadRequest, stream pb.Rpcftp_DownloadServer) error {
	log.Println("download rpc called: " + in.Path)
	file, err := os.Open(path.Join(s.rootDir, in.Path))
	if err != nil {
		log.Println(err.Error())
		st := status.New(codes.Internal, "Error: "+err.Error())
		return st.Err()
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err.Error())
			st := status.New(codes.Internal, "Error: "+err.Error())
			return st.Err()
		}
		if err := stream.Send(&pb.Chunk{
			FileName: path.Join(s.rootDir, in.Path),
			Content:  buf[:n],
		}); err != nil {
			log.Println(err.Error())
			st := status.New(codes.Internal, "Error: "+err.Error())
			return st.Err()
		}
	}
	return nil
}

func uploadErrHandle(text string, stream pb.Rpcftp_UploadServer, err error) {
	log.Println(text, err.Error())
	stream.SendAndClose(&pb.UploadStatus{
		Message: text + ": " + err.Error(),
		Code:    pb.UploadStatusCode_Failed,
	})
}

// Serve starts rpc server
func Serve(port, localPath string) error {
	listener, err := net.Listen("tcp", ":"+port)
	log.Println("creating listener on " + port)
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return err
	}
	srv := grpc.NewServer()
	log.Println("creating grpc server")
	pb.RegisterRpcftpServer(srv, NewServer(port, localPath))
	log.Printf("serving %v\n", localPath)
	if err := srv.Serve(listener); err != nil {
		log.Printf("failed to serve: %v\n", err)
		return err
	}
	return nil
}

package check_on_use

import (
	"CachingServer/pb"
	"context"
	"fmt"
	"io"
	"os"
)

type Check_on_use_server struct {
	pb.UnimplementedCheckOnUseServerServer
	AllFiles map[string]string
}

func NewServer() *Check_on_use_server {
	return &Check_on_use_server{AllFiles: make(map[string]string)}
}

func (s *Check_on_use_server) IsUpToDate(context context.Context, request *pb.CheckFileRequest) (*pb.CheckFileResponse, error) {
	// fmt.Println("Received a request!")
	checksum, ok := s.AllFiles[request.FileName]
	if !ok {
		fmt.Println("File doesn't exist")
		return &pb.CheckFileResponse{IsNewest: true}, nil
	}

	isNewest := checksum == request.VersionNumber
	return &pb.CheckFileResponse{IsNewest: isNewest}, nil
}

func (s *Check_on_use_server) PullFile(context context.Context, request *pb.PullFileRequest) (*pb.PullFileResponse, error) {

	// Make sure that the file exists
	_, err := os.Stat(request.FileName)
	if err != nil {
		return nil, err
	}

	// Make sure that the file is in the in-memory map
	checksum, ok := s.AllFiles[request.FileName]
	if !ok {
		return nil, fmt.Errorf("ENOENT")
	}

	// Load the file to memory
	fd, _ := os.Open(request.FileName)
	data, err := io.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	// Send the response
	return &pb.PullFileResponse{
		FileName:      request.FileName,
		VersionNumber: checksum,
		Data:          data,
	}, nil
}

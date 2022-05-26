package check_on_use

import (
	"CachingServer/pb"
	"context"
	id "github.com/google/uuid"
)

type Check_on_use_server struct {
	pb.UnimplementedCheckOnUseServerServer
	AllFiles map[string]id.UUID
}

func NewServer() *Check_on_use_server {
	return &Check_on_use_server{AllFiles: make(map[string]id.UUID)}
}

func (s *Check_on_use_server) IsUpToDate(context context.Context, request *pb.CheckFileRequest) (*pb.CheckFileResponse, error) {
	uuid, ok := s.AllFiles[request.FileName]
	if !ok {
		return &pb.CheckFileResponse{IsNewest: true}, nil
	}

	isNewest := uuid.String() == request.VersionNumber
	return &pb.CheckFileResponse{IsNewest: isNewest}, nil
}

func (s *Check_on_use_server) PullFile(context context.Context, request *pb.PullFileRequest) (*pb.PullFileResponse, error) {
	return &pb.PullFileResponse{
		FileName:      "NOTIMPLEMENTED",
		VersionNumber: "0",
		Data:          nil,
	}, nil
}

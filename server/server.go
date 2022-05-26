package main

import (
	"CachingServer/pb"
	"CachingServer/server_v1"
	"CachingServer/utils"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"sync"
	"time"
)

const (
	rootDir       string        = "./server_files/"
	checkInterval time.Duration = 5
)

var (
	wg sync.WaitGroup
)

func update(srv *server_v1.Check_on_use_server) {
	for range time.Tick(time.Second * checkInterval) {
		updateFileTable(srv)
	}
}
func updateFileTable(srv *server_v1.Check_on_use_server) {
	fd, err := os.Open(rootDir)
	if err != nil {
		fmt.Println(err)
	}
	defer fd.Close()

	// Get all file names in the directory
	list, err := fd.Readdirnames(-1)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range list {
		res, _ := utils.Md5sum(rootDir + f)
		srv.AllFiles[f] = res
	}
}

func main() {

	// Set up the gRPC server
	grpcServer := grpc.NewServer()
	srv := server_v1.NewServer(rootDir)
	pb.RegisterCheckOnUseServerServer(grpcServer, srv)
	listen, err := net.Listen("tcp", ":15213")
	if err != nil {
		fmt.Println(err)
	}

	// Create a go routine that updates the file table periodically
	go update(srv)

	err = grpcServer.Serve(listen)
	if err != nil {
		return
	}

	// Never exit the program
	wg.Add(1)
	wg.Wait()

}

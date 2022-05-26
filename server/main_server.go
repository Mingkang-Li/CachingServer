package main

import (
	"CachingServer/check_on_use"
	"CachingServer/pb"
	"CachingServer/utils"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"sync"
	"time"
)

const (
	rootDir       string        = "./files/"
	checkInterval time.Duration = 5
)

var (
	wg sync.WaitGroup
)

func update(srv *check_on_use.Check_on_use_server) {
	for range time.Tick(time.Second * checkInterval) {
		updateFileTable(srv)
	}
}
func updateFileTable(srv *check_on_use.Check_on_use_server) {
	fmt.Println("Table is updated!")
	fmt.Println(srv.AllFiles)
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
	grpcServer := grpc.NewServer()
	srv := check_on_use.NewServer()
	pb.RegisterCheckOnUseServerServer(grpcServer, srv)
	listen, err := net.Listen("tcp", ":15213")
	if err != nil {
		fmt.Println(err)
	}

	go update(srv)

	fmt.Println("Server is up and running")
	err = grpcServer.Serve(listen)
	if err != nil {
		return
	}

	// Never exit the program
	wg.Add(1)
	wg.Wait()

}

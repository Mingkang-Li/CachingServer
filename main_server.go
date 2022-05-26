package main

import (
	"CachingServer/check_on_use"
	"CachingServer/pb"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"
)

const (
	rootDir string = "./files"
)

func temp() {
	fd, err := os.Open(rootDir)
	if err != nil {
		fmt.Println(err)
	}

	prevList, err := fd.Readdirnames(-1)

	if err != nil {
		fmt.Println(err)
	}

	for {
		fd, _ = os.Open(rootDir)
		list, err := fd.Readdirnames(-1)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(list)
		if !cmp.Equal(list, prevList) {
			fmt.Println("Files changed!")
		}

		prevList = list
	}
	// Close the file!

}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterCheckOnUseServerServer(grpcServer, check_on_use.NewServer())
	listen, err := net.Listen("tcp", ":15213")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server is up and running")
	err = grpcServer.Serve(listen)
	if err != nil {
		return
	}

}

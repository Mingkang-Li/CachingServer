package main

import (
	"CachingServer/pb"
	"CachingServer/utils"
	"bytes"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

const (
	rootDir       string        = "../client_files/"
	checkInterval time.Duration = 5
)

var (
	wg sync.WaitGroup
)

func update(conn pb.CheckOnUseServerClient) {
	for range time.Tick(time.Second * checkInterval) {
		updateFileTable(conn)
	}
}
func updateFileTable(conn pb.CheckOnUseServerClient) {
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
		localChecksum, _ := utils.Md5sum(rootDir + f)
		res, _ := conn.IsUpToDate(context.Background(), &pb.CheckFileRequest{
			FileName:      f,
			VersionNumber: localChecksum,
		})

		// Update file if it is not up-to-date
		if !res.IsNewest {
			fmt.Println("Version is not newest for file", f)
			res, err := conn.PullFile(context.Background(), &pb.PullFileRequest{FileName: f})
			if err != nil {
				fmt.Println(err)
				continue
			}
			_ = os.Remove(rootDir + f)
			dstFd, err := os.OpenFile(rootDir+f, os.O_RDWR|os.O_CREATE, 0644)

			if err != nil {
				fmt.Println(err)
				continue
			}
			srcFd := bytes.NewReader(res.Data)

			io.Copy(dstFd, srcFd)
			dstFd.Close()
		}
	}
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:15213", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewCheckOnUseServerClient(conn)
	go update(client)

	// Never exit the main program
	wg.Add(1)
	wg.Wait()

}

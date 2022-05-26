package main

import (
	"CachingServer/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:15213", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCheckOnUseServerClient(conn)

	res, err := client.IsUpToDate(context.Background(), &pb.CheckFileRequest{
		FileName:      "What",
		VersionNumber: "No",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.IsNewest)

}

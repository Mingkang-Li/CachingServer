gen:
	protoc --go_out=. --go-grpc_out=. --proto_path=./proto proto/*.proto

clean:
	rm pb/*.go

run :
	go run ./server/server.go &
	go run ./client/client.go &
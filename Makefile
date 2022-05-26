gen:
	protoc --go_out=. --go-grpc_out=. --proto_path=./proto proto/*.proto

clean:
	rm pb/*.go

run server:
	go run main_server.go

run client:
	go run main_client.go

test:
	go test -race -cover ./tests/*.go
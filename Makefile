gen:
	protoc --go_out=. --go-grpc_out=. --proto_path=./proto proto/*.proto

clean:
	rm pb/*.go

run:
	go run main.go

test:
	go test -race -cover ./tests/*.go
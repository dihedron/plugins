.PHONY: all
all: driver grpc-plugin netrpc-plugin

.PHONY: driver
driver: main.go shared/*.go proto/*.go
	go build -o kv

.PHONY: grpc-plugin
grpc-plugin: plugin-go-grpc/main.go shared/*.go proto/*.go
	go build -o kv-go-grpc ./plugin-go-grpc

.PHONY: netrpc-plugin
netrpc-plugin: plugin-go-netrpc/main.go shared/*.go proto/*.go
	go build -o kv-go-netrpc ./plugin-go-netrpc

.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/kv.proto	

.PHONY: clean
clean:
	rm -rf kv kv_* kv-*
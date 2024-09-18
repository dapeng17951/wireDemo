build:
	protoc --proto_path=. --go_out=. *.proto
	protoc --proto_path=. --go-grpc_out=. kv.proto

clean:
	-rm -rf "../kv/*.*"
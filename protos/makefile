build:makedir
	protoc --proto_path=. --go_out=. *.proto
	protoc --proto_path=. --go-grpc_out=. kv.proto
	protoc --proto_path=. --csharp_out=kv_csharp --csharp_opt=base_namespace=KV *.proto
	
makedir:
	-mkdir kv_csharp

clean:
	-rm -rf "../kv/*.*"
	-rm -rf "../kv_csharp/*.*"
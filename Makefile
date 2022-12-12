

proto:
	protoc -I/usr/local/include \
		-Ipb \
		--go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		pb/*.proto
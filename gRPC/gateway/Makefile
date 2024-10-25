gen-cal:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out=repeated_path_param_separator=ssv:../ gatewaypb/gate.proto
run-server:
	go run main.go

run-proxy:
	go run proxy/proxy.go
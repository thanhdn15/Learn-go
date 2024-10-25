gen-cal:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculator/calculatorpb/calculator.proto

run-server:
	go run calculator/server/server.go

run-client:
	go run calculator/client/client.go

gen-mic:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative microservice/micropb/repo.proto

run-send-server:
	go run microservice/serviceSendMessage/main.go

run-receive-server:
	go run microservice/serviceReceiveMessage/main.go

build:
		protoc -I. --go_out=. --micro_out=. \
			proto/consignment/consignment.proto
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
		docker build -t shippy-service-consignment .
run:
		docker run -it -d -p 50051:50051  \
			-e MICRO_SERVER_ADDRESS=:50051 \
			shippy-service-consignment
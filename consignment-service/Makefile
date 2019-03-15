build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/shipper/consignment-service proto/consignment/consignment.proto --micro_out=.
run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service
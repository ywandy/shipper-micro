build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/shipper/vessel-service proto/vessel/vessel.proto --micro_out=.
run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service
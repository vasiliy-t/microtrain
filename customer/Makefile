GOPATH=/go/src
CWD = /go/src/github.com/vasiliy-t/microtrain/customer

proto/customer.pb.go: proto/customer.proto
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(GOPATH) \
		grpc/go:1.0 protoc -I$(CWD) --go_out=plugins=grpc:. $(CWD)/proto/customer.proto

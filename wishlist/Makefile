GOPATH=/go/src
CWD=$(GOPATH)/github.com/vasiliy-t/microtrain/wishlist

proto/wishlist.pb.go: proto/wishlist.proto
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(GOPATH) \
		grpc/go:1.0 protoc -I$(CWD) --go_out=. $(CWD)/proto/wishlist.proto 

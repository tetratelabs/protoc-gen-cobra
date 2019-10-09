PROTO_SOURCES = $(wildcard *.proto)
GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
COBRA_SOURCES = $(patsubst %.proto,%.cobra.pb.go,$(PROTO_SOURCES))

all: $(GO_SOURCES) $(COBRA_SOURCES)

$(GO_SOURCES) $(COBRA_SOURCES):
	protoc \
		-I. \
		--gofast_out=plugins=grpc:. \
		--cobra_out=plugins=client:. \
		$(PROTO_SOURCES)
	goimports -w $(COBRA_SOURCES)

clean:
	rm -f $(GO_SOURCES) $(COBRA_SOURCES)

deps:
	go get github.com/gogo/protobuf/protoc-gen-gofast
	go get github.com/tetratelabs/protoc-gen-cobra

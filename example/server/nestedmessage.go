package main

import (
	"golang.org/x/net/context"

	"github.com/tetratelabs/protoc-gen-cobra/example/pb"
)

type NestedMessage struct{}

var _ pb.NestedMessagesServer = NestedMessage{}

func (NestedMessage) Get(_ context.Context, req *pb.NestedRequest) (*pb.NestedResponse, error) {
	return &pb.NestedResponse{
		Return: req.Inner.Value + req.TopLevel.Value,
	}, nil
}

func (NestedMessage) GetDeeplyNested(_ context.Context, req *pb.DeeplyNested) (*pb.NestedResponse, error) {
	return &pb.NestedResponse{
		Return: req.L0.L1.L2.L3,
	}, nil
}

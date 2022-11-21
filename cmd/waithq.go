package main

import (
	"context"
	"log"
	"net"
	"strings"

	"github.com/CSC354/wathiq/proto"
	"google.golang.org/grpc"
)

const MOCKVALIDATOR = "LJeUtqLoLfFW4w"

type Waithq struct {
	proto.UnimplementedWathiqServer
}

func (*Waithq) Validate(ctx context.Context, in *proto.ValidateRequest) (response *proto.ValidateResponse, err error) {
	if strings.HasPrefix(in.Token, MOCKVALIDATOR) {
		response.Valid = true
	}
	return
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterWathiqServer(grpcServer, &Waithq{})
	go print("listening..")
	grpcServer.Serve(lis)
}

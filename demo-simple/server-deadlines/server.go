package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"

	pb "go-grpc-example/demo-simple/proto"

	"google.golang.org/grpc"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {

	for i := 0; i < 6; i++ {
		fmt.Println(i)
		if ctx.Err() == context.Canceled {
			log.Println(ctx.Err())
			return nil, status.Errorf(codes.Canceled, "SearchService.Search canceled")
		}
		time.Sleep(2 * time.Second)
	}
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/coooller/chat-server/pkg/chat_v1"
)

const grpcPort = 50052

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) Create(_ context.Context, in *desc.CreateChatIn) (*desc.CreateChatOut, error) {
	log.Printf("Create Chat: %v", in)

	return &desc.CreateChatOut{Id: gofakeit.Int64()}, nil
}

func (s *server) Delete(_ context.Context, in *desc.DeleteChatIn) (*emptypb.Empty, error) {
	log.Printf("Delete Chat: %v", in)
	return nil, nil
}

func (s *server) SendMessage(_ context.Context, in *desc.SendMessageIn) (*emptypb.Empty, error) {
	log.Printf("Send Message: %v", in)
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

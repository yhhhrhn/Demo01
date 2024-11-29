package grpc

import (
	pb "awesomeProject/grpc/proto"
	"awesomeProject/service"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedTaskServer
}

func (s *server) GetTask(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	//res, err := poker.PokerEvaluator(req.Hand, req.River)
	fmt.Println(req.GetId())
	var task = service.GetTaskById(req.GetId())
	return &pb.TaskResponse{
		Id:          task.Id,
		Description: task.Description,
		Status:      task.Status,
	}, nil
}
func InitGrpc() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTaskServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

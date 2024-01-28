package main

import (
    "context"
    "log"
    "net"
    "google.golang.org/grpc"
    pb "example.com/myproject/hello" // 更新為您的模組路徑
)

// server 用於實現 hello.proto 中定義的 GreeterServer。
type server struct {
    pb.UnimplementedGreeterServer
}

// SayHello 實現了 hello.proto 中定義的 GreeterServer 接口。
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    // 實現 SayHello 業務邏輯
    return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

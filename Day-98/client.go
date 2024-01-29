package main

import (
    "context"
    "log"
    "time"
    "google.golang.org/grpc"
    pb "example.com/myproject/hello" // 您的 protobuf 包
)

func main() {
    // 設置服務器地址
    const address = "localhost:50051"

    // 建立連接
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("無法連接: %v", err)
    }
    defer conn.Close()
    c := pb.NewGreeterClient(conn)

    // 設置上下文超時
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // 呼叫 SayHello 方法
    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "My First GRPC"})
    if err != nil {
        log.Fatalf("無法調用 SayHello: %v", err)
    }
    log.Printf("來自 Server 的響應: %s", r.GetMessage())
}

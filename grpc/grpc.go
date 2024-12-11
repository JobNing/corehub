package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func RegisterGRPC(port int64, hand func(s *grpc.Server)) error {
	fmt.Printf("建立tcp监听\n")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	fmt.Printf("初始化grpc\n")
	s := grpc.NewServer()

	fmt.Printf("开始执行闭包函数\n")

	hand(s)

	fmt.Printf("开启一个映射\n")
	reflection.Register(s)
	fmt.Printf("grpc server is started listening on port %d \n", port)
	return s.Serve(lis)
}

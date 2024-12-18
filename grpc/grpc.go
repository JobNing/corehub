package grpc

import (
	"fmt"
	"github.com/JobNing/corehub/config"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

//func getHostIP() (string, error) {
//	addrs, err := net.InterfaceAddrs()
//
//	if err != nil {
//		return "", err
//	}
//
//	for _, address := range addrs {
//		// 检查ip地址判断是否回环地址
//		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
//			if ipnet.IP.To4() != nil {
//				return ipnet.IP.String(), nil
//			}
//		}
//	}
//	return "", fmt.Errorf("ip 获取失败")
//}

func RegisterGRPC(port int64, hand func(s *grpc.Server)) error {
	// TODO:自动获取ip
	//ip := "127.0.0.1"
	err := config.RegisterServiceInstance(viper.GetString("ip"), port, viper.GetString("service-name"))
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	hand(s)

	fmt.Printf("开启一个映射\n")
	reflection.Register(s)
	fmt.Printf("grpc server is started listening on port %d \n", port)
	return s.Serve(lis)
}

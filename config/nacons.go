package config

import (
	"github.com/spf13/viper"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func getConfig() vo.NacosClientParam {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(viper.GetString("n-ip"),
			viper.GetUint64("n-port"),
			constant.WithContextPath(viper.GetString("n-path"))),
	}
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(viper.GetString("namespace")),
		constant.WithTimeoutMs(viper.GetUint64("timeout-ms")),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(viper.GetString("log-dir")),
		constant.WithCacheDir(viper.GetString("cache-dir")),
		constant.WithLogLevel(viper.GetString("log-level")),
	)
	return vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	}
}

func GetConfig() (string, error) {
	client, err := clients.NewConfigClient(getConfig())
	if err != nil {
		return "", err
	}

	return client.GetConfig(vo.ConfigParam{
		DataId: viper.GetString("data-id"),
		Group:  viper.GetString("group"),
	})
}

func RegisterServiceInstance(ip string, port int64, serviceName string) error {
	//获取服务注册与发现的Client
	client, err := clients.NewNamingClient(getConfig())
	if err != nil {
		return err
	}

	//通过Client调用服务注册的方法，填充ip 端口和服务名
	_, err = client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        uint64(port),
		ServiceName: serviceName,
	})
	return err
}

//
//TODO:监听配置
//func ListenConfig() {
//	//Listen config change,key=dataId+group+namespaceId.
//	err = client.ListenConfig(vo.ConfigParam{
//		DataId: "test-data",
//		Group:  "test-group",
//		OnChange: func(namespace, group, dataId, data string) {
//			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
//		},
//	})
//}

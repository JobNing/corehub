package config

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/spf13/viper"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func getClient() (config_client.IConfigClient, error) {
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(viper.GetString("n-ip"),
			viper.GetUint64("n-port"),
			constant.WithContextPath(viper.GetString("n-path"))),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(viper.GetString("namespace")),
		constant.WithTimeoutMs(viper.GetUint64("timeout-ms")),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(viper.GetString("log-dir")),
		constant.WithCacheDir(viper.GetString("cache-dir")),
		constant.WithLogLevel(viper.GetString("log-level")),
	)

	// create config client
	return clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
}

func GetConfig() (string, error) {
	client, err := getClient()
	if err != nil {
		return "", err
	}

	fmt.Println(viper.GetString("data-id"))
	fmt.Println(viper.GetString("group"))
	return client.GetConfig(vo.ConfigParam{
		DataId: viper.GetString("data-id"),
		Group:  viper.GetString("group"),
	})
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

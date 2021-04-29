package nacos_viper_remote

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/common/logger"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"strings"
)

type nacosConfigManager struct {
	client config_client.IConfigClient
	option *Option
}

func NewNacosConfigManager(option *Option) (*nacosConfigManager, error) {
	var serverConfigs []constant.ServerConfig
	urls := strings.Split(option.Url, ";")
	for _, url := range urls {
		serverConfigs = append(serverConfigs, constant.ServerConfig{
			ContextPath: "/nacos",
			IpAddr:      url,
			Port:        option.Port,
		})
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:         option.NamespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "info",
	}
	if option.Auth != nil && option.Auth.Enable {
		clientConfig.Username = option.Auth.User
		clientConfig.Password = option.Auth.Password
	}
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	manager := &nacosConfigManager{client: client, option: option}

	return manager, err
}

func (cm *nacosConfigManager) Get(dataId string) ([]byte, error) {
	//get config
	content, err := cm.client.GetConfig(vo.ConfigParam{
		DataId: cm.option.Config.DataId,
		Group:  cm.option.GroupName,
	})
	return []byte(content), err
}

func (cm *nacosConfigManager) Watch(dataId string, stop chan bool) <-chan *viper.RemoteResponse {
	resp := make(chan *viper.RemoteResponse)

	configParams := vo.ConfigParam{
		DataId: cm.option.Config.DataId,
		Group:  cm.option.GroupName,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId)
			resp <- &viper.RemoteResponse{
				Value: []byte(data),
				Error: nil,
			}
		},
	}
	err := cm.client.ListenConfig(configParams)
	if err != nil {
		return nil
	}

	go func() {
		for {
			select {
			case <-stop:
				_ = cm.client.CancelListenConfig(configParams)
				return
			}
		}
	}()

	return resp
}

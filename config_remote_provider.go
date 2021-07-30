package nacos_viper_remote

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
)

type ViperRemoteProvider struct {
	configType string
	configSet  string
}

func NewRemoteProvider(configType string) *ViperRemoteProvider {
	return &ViperRemoteProvider{
		configType: configType,
		configSet:  "yoyogo.cloud.discovery.metadata"}
}

func (provider *ViperRemoteProvider) GetProvider(runtimeViper *viper.Viper) *viper.Viper {
	var option *Option
	err := runtimeViper.Sub(provider.configSet).Unmarshal(&option)
	if err != nil {
		panic(err)
		return nil
	}
	SetOptions(option)
	remote_viper := viper.New()
	err = remote_viper.AddRemoteProvider("nacos", "localhost", "")
	if provider.configType == "" {
		provider.configType = "yaml"
	}
	remote_viper.SetConfigType(provider.configType)
	err = remote_viper.ReadRemoteConfig()
	if err == nil {
		//err = remote_viper.WatchRemoteConfigOnChannel()
		if err == nil {
			fmt.Println("config center ..........")
			fmt.Println("used remote viper by Nacos")
			fmt.Printf("Nacos config: namespace: %s , group: %s", option.NamespaceId, option.GroupName)
			return remote_viper
		}
	} else {
		panic(err)
	}
	return runtimeViper
}

func (provider *ViperRemoteProvider) WatchRemoteConfigOnChannel(remoteViper *viper.Viper) <-chan bool {
	updater := make(chan bool)

	respChan, _ := viper.RemoteConfig.WatchChannel(DefaultRemoteProvider())
	go func(rc <-chan *viper.RemoteResponse) {
		for {
			b := <-rc
			reader := bytes.NewReader(b.Value)
			_ = remoteViper.ReadConfig(reader)
			// configuration on changed
			updater <- true
		}
	}(respChan)

	return updater
}

package nacos_viper_remote

import (
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

func (provider *ViperRemoteProvider) GetProvider(runtime_viper *viper.Viper) *viper.Viper {
	var option *Option
	err := runtime_viper.Sub(provider.configSet).Unmarshal(&option)
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
		err = remote_viper.WatchRemoteConfigOnChannel()
		if err == nil {
			fmt.Println("used remote viper")
			return remote_viper
		}
	} else {
		panic(err)
	}
	return runtime_viper
}

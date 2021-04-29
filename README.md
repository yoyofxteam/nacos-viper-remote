# Viper remote for Nacos

Golang configuration,use to Viper reading from remote Nacos config systems. Viper remote for Naocs.

```go
runtime_viper := viper.New()

remote.SetOptions(&remote.Option{
   Url:         "localhost",
   Port:        80,
   NamespaceId: "public",
   GroupName:   "DEFAULT_GROUP",
   Config: 	remote.Config{ DataId: "config_dev" },
   Auth:        nil,
})

err := remote_viper.AddRemoteProvider("nacos", "localhost", "")
remote_viper.SetConfigType("yaml")

_ = remote_viper.ReadRemoteConfig()             //sync get remote configs to remote_viper instance memory . for example , remote_viper.GetString(key)
_ = remote_viper.WatchRemoteConfigOnChannel()   //async watch , auto refresh configs.

appName := remote_viper.GetString("key")   // sync get config by key

fmt.Println(appName)
```

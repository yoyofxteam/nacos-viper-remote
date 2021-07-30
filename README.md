# Viper remote for Nacos

Golang configuration,use to Viper reading from remote Nacos config systems. Viper remote for Naocs.

```go
config_viper := viper.New()

remote.SetOptions(&remote.Option{
   Url:         "localhost",
   Port:        80,
   NamespaceId: "public",
   GroupName:   "DEFAULT_GROUP",
   Config: 	remote.Config{ DataId: "config_dev" },
   Auth:        nil,
})

remote_viper := viper.New()
err := remote_viper.AddRemoteProvider("nacos", "localhost", "")
remote_viper.SetConfigType("yaml")
err = remote_viper.ReadRemoteConfig()    //sync get remote configs to remote_viper instance memory . for example , remote_viper.GetString(key)

if err == nil {
    config_viper = remote_viper
    fmt.Println("used remote viper")
    provider := remote.NewRemoteProvider("yaml")
    respChan := provider.WatchRemoteConfigOnChannel(config_viper)

    go func(rc <-chan bool) {
        for {
            <-rc
            fmt.Printf("remote async: %s", config_viper.GetString("yoyogo.application.name"))
        }
    }(respChan)
}

go func() {
    for {
        time.Sleep(time.Second * 30) // delay after each request
        appName = config_viper.GetString("yoyogo.application.name")
        fmt.Println("sync:" + appName)
    }
}()
```

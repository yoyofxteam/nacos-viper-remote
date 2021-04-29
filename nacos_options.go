package nacos_viper_remote

type Option struct {
	Url         string `mapstructure:"url"`
	Port        uint64 `mapstructure:"port"`
	NamespaceId string `mapstructure:"namespace"`
	GroupName   string `mapstructure:"group"`
	Config      Config `mapstructure:"configserver"`
	Auth        *Auth  `mapstructure:"auth"`
}

type Config struct {
	DataId string `mapstructure:"dataId"`
}

type Auth struct {
	Enable   bool   `mapstructure:"enable"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

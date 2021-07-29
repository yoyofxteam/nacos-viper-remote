package nacos_viper_remote

type nacosRemoteProvider struct {
	provider      string
	endpoint      string
	path          string
	secretKeyring string
}

func DefaultRemoteProvider() *nacosRemoteProvider {
	return &nacosRemoteProvider{provider: "nacos", endpoint: "localhost", path: "", secretKeyring: ""}
}

func (rp nacosRemoteProvider) Provider() string {
	return rp.provider
}

func (rp nacosRemoteProvider) Endpoint() string {
	return rp.endpoint
}

func (rp nacosRemoteProvider) Path() string {
	return rp.path
}

func (rp nacosRemoteProvider) SecretKeyring() string {
	return rp.secretKeyring
}

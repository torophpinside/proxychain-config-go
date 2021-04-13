package loader_proxy

type ProxyLoaderInterface interface {
	Load() ([]string, error)
}

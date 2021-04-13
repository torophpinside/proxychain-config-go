package loader_proxy_proxy_source

type ProxySourceInterface interface {
	Get() (string, error)
}

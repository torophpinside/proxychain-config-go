package loader_proxy

import (
	"proxychain-config-go/helper"
	loader_proxy_proxy_source "proxychain-config-go/loader/proxy/proxy_source"
)

type randomProxyLoader struct {
	source []loader_proxy_proxy_source.ProxySourceInterface
}

func NewRandomProxyLoader(s []loader_proxy_proxy_source.ProxySourceInterface) ProxyLoaderInterface {
	return &randomProxyLoader{
		source: s,
	}
}

func (r *randomProxyLoader) Load() ([]string, error) {
	var l []string
	md := len(r.source)
	for i := 0; i < 50; i++ {
		s := r.source[i%md]
		p, err := s.Get()
		if err != nil {
			return nil, err
		}
		if p == "://" {
			continue
		}

		l = append(l, p)
	}

	l, err := helper.RawProxyChecker(l)
	if err != nil {
		return nil, err
	}

	return l, nil
}

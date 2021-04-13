package loader_proxy

import (
	"net/http"
	"net/url"
	loader_proxy_proxy_source "proxychain-config-go/loader/proxy/proxy_source"
	"time"
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
	for i := 0; i < 50; i++ {
		s := r.source[i%3]
		p, err := s.Get()
		if err != nil {
			return nil, err
		}
		if p == "://" {
			continue
		}

		proxyURL, err := url.Parse(p)
		if err != nil {
			return nil, err
		}

		myClient := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
			Timeout: 15 * time.Second,
		}
		resp, err := myClient.Get("https://www.google.com/")
		if err != nil || resp == nil {
			time.Sleep(5 * time.Second)
			continue
		}

		l = append(l, p)
	}

	return l, nil
}

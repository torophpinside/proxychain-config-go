package loader_proxy

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_configConnection "proxychain-config-go/config/connection"
	"proxychain-config-go/helper"
)

type freeProxyListProxyLoader struct {
	conn _configConnection.ConnectionInterface
}

func (f *freeProxyListProxyLoader) Load() ([]string, error) {
	body, err := f.conn.Get("https://free-proxy-list.net/")
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var fUrls []string
	doc.Find("#proxylisttable tbody tr").Each(func(i int, s *goquery.Selection) {
		ip := s.Find("td:first-child").Text()
		port := s.Find("td:nth-child(2)").Text()

		p := fmt.Sprintf("http://%s:%s", ip, port)
		fUrls = append(fUrls, p)
	})

	fUrls, err = helper.RawProxyChecker(fUrls)
	if err != nil {
		return nil, err
	}

	return fUrls, nil
}

func NewFreeProxyListProxyLoader(c _configConnection.ConnectionInterface) ProxyLoaderInterface {
	return &freeProxyListProxyLoader{
		conn: c,
	}
}

package loader_proxy

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	config_connection "proxychain-config-go/config/connection"
	"proxychain-config-go/helper"
)

type proxyListPlusProxyLoader struct {
	conn config_connection.ConnectionInterface
}

func NewProxyListPlusProxyLoader(c config_connection.ConnectionInterface) ProxyLoaderInterface {
	return &proxyListPlusProxyLoader{
		conn: c,
	}
}

func (apl *proxyListPlusProxyLoader) Load() ([]string, error) {
	body, err := apl.conn.Get("https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-1")
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var fUrls []string
	doc.Find("body div table:nth-child(2) tr").Each(func(i int, s *goquery.Selection) {
		ip := s.Find("td:first-child").Text()
		port := s.Find("td:nth-child(2)").Text()
		https := s.Find("td:nth-child(7)").Text()

		if https == "yes" {
			return
		}

		p := fmt.Sprintf("http://%s:%s", ip, port)
		fUrls = append(fUrls, p)
	})

	fUrls, err = helper.RawProxyChecker(fUrls)
	if err != nil {
		return nil, err
	}

	return fUrls, nil
}

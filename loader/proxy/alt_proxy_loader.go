package loader_proxy

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	config_connection "proxychain-config-go/config/connection"
	"proxychain-config-go/helper"
)

type altProxyLoader struct {
	conn config_connection.ConnectionInterface
}

func NewAltProxyLoader(c config_connection.ConnectionInterface) ProxyLoaderInterface {
	return &altProxyLoader{
		conn: c,
	}
}

func (apl *altProxyLoader) Load() ([]string, error) {
	body, err := apl.conn.Get("https://free-proxy-list.net/")
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var fUrls []string
	doc.Find("#list .container .table-responsive tbody tr").Each(func(i int, s *goquery.Selection) {
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

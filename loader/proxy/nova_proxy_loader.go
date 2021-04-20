package loader_proxy

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	config_connection "proxychain-config-go/config/connection"
	"proxychain-config-go/helper"
	"strings"
)

type novaProxyLoader struct {
	conn config_connection.ConnectionInterface
}

func NewNovaProxyLoader(c config_connection.ConnectionInterface) ProxyLoaderInterface {
	return &novaProxyLoader{
		conn: c,
	}
}

func (npl *novaProxyLoader) Load() ([]string, error) {
	body, err := npl.conn.Get("https://www.proxynova.com/proxy-server-list/")
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var fUrls []string
	doc.Find("#tbl_proxy_list tbody tr").Each(func(i int, s *goquery.Selection) {
		ipRaw := s.Find("td:first-child abbr script").Text()
		portRaw := s.Find("td:nth-child(2)").Text()

		portRaw = strings.Replace(portRaw, "\n", "", -1)
		port := strings.Replace(portRaw, " ", "", -1)

		ipRaw = strings.Replace(ipRaw, "document.write('", "", -1)
		ip := strings.Replace(ipRaw, "');", "", -1)

		p := fmt.Sprintf("http://%s:%s", ip, port)
		if p == "http://:" {
			return
		}
		fUrls = append(fUrls, p)
	})

	fUrls, err = helper.RawProxyChecker(fUrls)
	if err != nil {
		return nil, err
	}

	return fUrls, nil
}

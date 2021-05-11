package loader_proxy

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_configConnection "proxychain-config-go/config/connection"
	"proxychain-config-go/helper"
	"strings"
)

type spysoneProxyLoader struct {
	conn _configConnection.ConnectionInterface
}

func (f *spysoneProxyLoader) Load() ([]string, error) {
	body, err := f.conn.Get("https://spys.one/en/")
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var fUrls []string
	doc.Find("table:nth-child(3) table:nth-child(1) table:nth-child(2) tbody tr").Each(func(i int, s *goquery.Selection) {
		if i < 2 {
			return
		}
		ipRaw := s.Find("td:first-child font").Text()
		ip := strings.Split(ipRaw, "document")

		t := s.Find("td:nth-child(2) font").Text()
		ipPorts := f.withPorts(ip[0], t)
		for _, v := range ipPorts {
			fUrls = append(fUrls, v)
		}
	})

	fUrls, err = helper.RawProxyChecker(fUrls)
	if err != nil {
		return nil, err
	}

	return fUrls, nil
}

func (f *spysoneProxyLoader) withPorts(ipstr string, tc string) []string {
	ports := []string{
		"999",
		"1080",
		"3128",
		"4145",
		"6667", "6699",
		"8080", "8081", "8888",
		"18080",
		"20183",
		"38473",
		"41419", "44239",
		"53281", "51008", "55443",
	}
	var ipPort []string
	var ptc string
	if strings.Contains(tc, "HTTP") || strings.Contains(tc, "HTTPS") {
		ptc = "http"
	}
	if tc == "SOCKS5" {
		ptc = "socks5"
	}
	for _, v := range ports {
		ipPort = append(ipPort, fmt.Sprintf("%s://%s:%s", ptc, ipstr, v))
	}

	return ipPort
}

func NewSpysoneProxyLoader(c _configConnection.ConnectionInterface) ProxyLoaderInterface {
	return &spysoneProxyLoader{
		conn: c,
	}
}

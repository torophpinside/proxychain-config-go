package loader_proxy

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	config_connection "proxychain-config-go/config/connection"
	"strings"
	"sync"
	"time"
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
		log.Fatal(err)
	}

	r := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
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

	var urls []string
	var wg sync.WaitGroup
	for _, p := range fUrls {
		wg.Add(1)
		go func(p string, urls *[]string, wg *sync.WaitGroup) {
			defer wg.Done()
			proxyURL, err := url.Parse(p)
			if err != nil {
				fmt.Println("")
				log.Fatalf("Failed to parse proxy URL: %v\n", err)
			}

			myClient := http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyURL(proxyURL),
				},
				Timeout: 20 * time.Second,
			}
			resp, err := myClient.Get("https://www.bhphotovideo.com/")
			if err != nil || resp == nil {
				fmt.Print(".")
				return
			}
			fmt.Print("|")
			*urls = append(*urls, p)
		}(p, &urls, &wg)
	}
	wg.Wait()

	fmt.Println("")
	return urls, nil
}

package loader_proxy

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	config_connection "proxychain-config-go/config/connection"
	"sync"
	"time"
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
		log.Fatal(err)
	}

	r := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
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

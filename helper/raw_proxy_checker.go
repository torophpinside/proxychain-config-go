package helper

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func RawProxyChecker(fUrls []string) ([]string, error) {
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

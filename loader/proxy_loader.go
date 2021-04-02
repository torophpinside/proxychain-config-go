package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var clients []*http.Client
var randomClient *http.Client
var torClient *http.Client
var lastClientKey = 0
var randomEvenOdd = 0

func initTorClient() {
	up, err := url.Parse("socks5://127.0.0.1:9050")
	if err != nil {
		log.Fatalf("Failed to parse proxy URL: %v\n", err)
	}

	t := &http.Transport{Proxy: http.ProxyURL(up)}
	c := &http.Client{Transport: t}

	torClient = c
}

func pubProxy() string {
	type proxyItem struct {
		IPPort string `json:"ipPort"`
		Type   string `json:"type"`
	}

	type proxyData struct {
		Data  []proxyItem `json:"data"`
		Count int         `json:"count"`
	}

	pd := proxyData{}

	resp, err := GetHTTPTorClient().Get("http://pubproxy.com/api/proxy?user_agent=true&speed=3&type=http&last_check=600&format=json")
	if err != nil {
		log.Fatalf("Failed to issue GET request: %v\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the body: %v\n", err)
	}

	if strings.Contains(string(body), "No proxy") {
		log.Fatalln("no random proxy found (no proxy pubproxy)")
	}

	err = json.Unmarshal(body, &pd)
	if err != nil {
		log.Fatalf("Failed to read the body: %v\n", err)
	}

	if pd.Count == 0 {
		log.Println("no random proxy found (pubproxy)")
		return "://"
	}

	return fmt.Sprintf("%s://%s", pd.Data[0].Type, pd.Data[0].IPPort)
}

func getProxyListAPI() string {
	type proxyData struct {
		AllowsHTTPS bool   `json:"allowsHttps"`
		IP          string `json:"ip"`
		Port        int    `json:"port"`
	}

	pd := proxyData{}

	resp, err := GetHTTPTorClient().Get("https://api.getproxylist.com/proxy?minDownloadSpeed=300&maxSecondsToFirstByte=10&allowsPost=1&maxConnectTime=15&protocol[]=http&allowsUserAgentHeader=1&lastTested=800")
	if err != nil {
		log.Fatalf("Failed to issue GET request: %v\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the body: %v\n", err)
	}

	err = json.Unmarshal(body, &pd)
	if err != nil {
		log.Fatalf("Failed to read the body: %v\n", err)
	}

	if pd.IP == "" {
		log.Println("no random proxy found (getproxy)")
		return "://"
	}

	sec := "http"
	return fmt.Sprintf("%s://%s:%d", sec, pd.IP, pd.Port)
}

func gimmeProxy() string {
	type proxyData struct {
		IPPort string `json:"ipPort"`
		Type   string `json:"type"`
	}

	pd := proxyData{}

	resp, err := GetHTTPTorClient().Get("https://gimmeproxy.com/api/getProxy?post=true&anonymityLevel=1&protocol=http")
	if err != nil {
		log.Fatalf("Failed to issue GET request: %v\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the body: %v\n", err)
	}

	if strings.Contains(string(body), "No proxy") {
		log.Fatalln("no random proxy found (no proxy gimmeproxy)")
	}

	err = json.Unmarshal(body, &pd)
	if err != nil {
		log.Fatalf("Failed to read the body: %v\n", err)
	}

	return fmt.Sprintf("%s://%s", pd.Type, pd.IPPort)
}

// ListCrawlerProxy ...
func ListCrawlerProxy(max int) []string {
	res, err := http.Get("https://free-proxy-list.net/")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var urls []string
	doc.Find("#list .container .table-responsive tbody tr").Each(func(i int, s *goquery.Selection) {
		if len(urls) > max {
			return
		}

		ip := s.Find("td:first-child").Text()
		port := s.Find("td:nth-child(2)").Text()
		https := s.Find("td:nth-child(7)").Text()

		if https == "yes" {
			return
		}

		p := fmt.Sprintf("http://%s:%s", ip, port)

		proxyURL, err := url.Parse(p)
		if err != nil {
			fmt.Println("")
			log.Fatalf("Failed to parse proxy URL: %v\n", err)
		}

		myClient := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
			Timeout: 15 * time.Second,
		}
		resp, err := myClient.Get("https://www.bhphotovideo.com/")
		if err != nil || resp == nil {
			fmt.Print(".")
			return
		}
		fmt.Print("|")
		urls = append(urls, p)
	})
	fmt.Println("")
	return urls
}

// InitRandomProxy ...
func InitRandomProxy() string {
	var myClient *http.Client
	log.Println("finding proxy")
	var p string
	for {
		if randomEvenOdd == 0 {
			fmt.Print(".")
			p = pubProxy()
		}
		if randomEvenOdd == 1 {
			fmt.Print("..")
			p = gimmeProxy()
		}
		if randomEvenOdd == 2 {
			fmt.Print("...")
			p = getProxyListAPI()
		}
		randomEvenOdd++
		if randomEvenOdd > 2 {
			randomEvenOdd = 0
		}

		if p == "://" {
			continue
		}

		proxyURL, err := url.Parse(p)
		if err != nil {
			fmt.Println("")
			log.Fatalf("Failed to parse proxy URL: %v\n", err)
		}

		myClient = &http.Client{
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
		fmt.Println("")
		fmt.Println("Proxy found:", p)
		break
	}

	randomClient = myClient
	return p
}

// GetHTTPClient ...
func GetHTTPClient(opt string) *http.Client {
	if opt == "random" {
		return randomClient
	}

	client := clients[lastClientKey]
	lastClientKey++
	if lastClientKey >= len(clients) {
		lastClientKey = 0
	}
	return client
}

// GetHTTPTorClient ...
func GetHTTPTorClient() *http.Client {
	initTorClient()
	return torClient
}

// ListCrawlerProxyAlt ...
func ListCrawlerProxyAlt(max int) []string {
	res, err := http.Get("https://free-proxy-list.net/")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
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
	return urls
}

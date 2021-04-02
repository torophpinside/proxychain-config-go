package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"proxychain-config-go/loader"
	"strings"
)

type options struct {
	searchType string
	max        int
}

func main() {
	log.Println("start proxychains configuration")

	opt := options{}
	flag.StringVar(&opt.searchType, "search", "random", "Type of search proxylst: random/list")
	flag.IntVar(&opt.max, "max", 9, "Max of IPs proxies")
	flag.Parse()

	bs := "ZHluYW1pY19jaGFpbgpwcm94eV9kbnMgCnRjcF9yZWFkX3RpbWVfb3V0IDYwMDAwCnRjcF9jb25uZWN0X3RpbWVfb3V0IDMwMDAwCltQcm94eUxpc3RdCiN7UFJPWFk5fQoje1BST1hZOH0KI3tQUk9YWTd9CiN7UFJPWFk2fQoje1BST1hZNX0KI3tQUk9YWTR9CiN7UFJPWFkzfQoje1BST1hZMn0KI3tQUk9YWTF9"

	sDec, err := base64.StdEncoding.DecodeString(bs)
	if err != nil {
		log.Fatalln("error decoding tpl", err)
	}
	tpl := []byte(sDec)

	if opt.searchType == "list" {
		lp := loader.ListCrawlerProxyAlt(opt.max)
		if len(lp) == 0 {
			log.Fatalln("no proxy found")
		}
		for i := 1; i <= opt.max; i++ {
			lbl := fmt.Sprintf("#{PROXY%d}", i)
			tpl = bytes.Replace(tpl, []byte(lbl), []byte(formatConfig(lp[i-1])), -1)
		}

	} else {
		p := loader.InitRandomProxy()
		log.Println(p)
		tpl = bytes.Replace(tpl, []byte("#{PROXY1}"), []byte(formatConfig(p)), -1)

		p = loader.InitRandomProxy()
		log.Println(p)
		tpl = bytes.Replace(tpl, []byte("#{PROXY2}"), []byte(formatConfig(p)), -1)

		p = loader.InitRandomProxy()
		log.Println(p)
		tpl = bytes.Replace(tpl, []byte("#{PROXY3}"), []byte(formatConfig(p)), -1)

		p = loader.InitRandomProxy()
		log.Println(p)
		tpl = bytes.Replace(tpl, []byte("#{PROXY4}"), []byte(formatConfig(p)), -1)

		p = loader.InitRandomProxy()
		log.Println(p)
		tpl = bytes.Replace(tpl, []byte("#{PROXY5}"), []byte(formatConfig(p)), -1)

		p = loader.InitRandomProxy()
		log.Println(p)
		tpl = bytes.Replace(tpl, []byte("#{PROXY6}"), []byte(formatConfig(p)), -1)
	}

	if err := ioutil.WriteFile("/etc/proxychains.conf", tpl, 0666); err != nil {
		log.Fatalln(err)
	}

	log.Println("proxychains configured!")
}

func formatConfig(u string) string {
	spl := strings.Split(u, ":")
	ip := spl[1]
	ip = ip[2:len(spl[1])]
	return fmt.Sprintf("%s    %s    %s", spl[0], ip, spl[2])
}

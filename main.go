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

	bs := "ZHluYW1pY19jaGFpbgpwcm94eV9kbnMgCnRjcF9yZWFkX3RpbWVfb3V0IDE1MDAwCnRjcF9jb25uZWN0X3RpbWVfb3V0IDgwMDAKW1Byb3h5TGlzdF0KI3tQUk9YWTF9CiN7UFJPWFkyfQoje1BST1hZM30KI3tQUk9YWTR9CiN7UFJPWFk1fQoje1BST1hZNn0KI3tQUk9YWTd9CiN7UFJPWFk4fQoje1BST1hZOX0="

	sDec, err := base64.StdEncoding.DecodeString(bs)
	if err != nil {
		log.Fatalln("error decoding tpl", err)
	}
	tpl := []byte(sDec)

	if opt.searchType == "list" {
		lp := loader.ListCrawlerProxy(opt.max)
		log.Println(lp)
		for i := 1; i <= opt.max; i++ {
			lbl := fmt.Sprintf("#{PROXY%d}", i)
			tpl = bytes.Replace(tpl, []byte(lbl), []byte(formatConfig(lp[0])), -1)
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

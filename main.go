package main

import (
	"io/ioutil"
	"log"
	"proxychain-config-go/config"
	config_template "proxychain-config-go/config/template"
	"proxychain-config-go/interpreter"
)

func main() {
	log.Println("start proxychains configuration")

	opt := config.GetOptions()
	tpl, err := config_template.NewEncodedTemplate().GetTemplate()
	if err != nil {
		log.Fatalln("could not find template", err)
	}

	p := interpreter.Factory(opt, tpl)
	pTpl, err := p.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	if err := ioutil.WriteFile("/etc/proxychains.conf", pTpl, 0666); err != nil {
		log.Fatalln(err)
	}

	log.Println("proxychains configured!")
}

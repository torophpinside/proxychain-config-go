package main

import (
	"io/ioutil"
	"log"
	"proxychain-config-go/config"
	config_connection "proxychain-config-go/config/connection"
	config_template "proxychain-config-go/config/template"
	helper_formatter "proxychain-config-go/helper/formatter"
	"proxychain-config-go/interpreter"
)

func main() {
	log.Println("start proxychains configuration")

	opt := config.GetOptions()
	tc := config_connection.NewTorConnection()
	err := tc.Connect()
	if err != nil {
		log.Fatalln("could not start connection", err)
	}
	frm := helper_formatter.NewIPFormatterHelper()

	tpl, err := config_template.NewEncodedTemplate().GetTemplate()
	if err != nil {
		log.Fatalln("could not find template", err)
	}

	p := interpreter.Factory(opt, tpl, frm, tc)

	pTpl, err := p.Parse()
	if err != nil {
		log.Fatalln("could not parse template - ", err)
	}

	if err := ioutil.WriteFile("/etc/proxychains.conf", pTpl, 0666); err != nil {
		log.Fatalln(err)
	}

	log.Println("proxychains configured!")
}

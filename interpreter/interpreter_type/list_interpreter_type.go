package parser_parser_type

import (
	"bytes"
	"fmt"
	"log"
	"proxychain-config-go/config"
	helper_formatter "proxychain-config-go/helper/formatter"
	"proxychain-config-go/loader"
	"strings"
)

type listInterpreterType struct {
	opt         config.Options
	tpl         []byte
	ipFormatter helper_formatter.FormatterInterface
}

func NewListInterpreterType(opt config.Options, tpl []byte, f helper_formatter.FormatterInterface) InterpreterTypeInterface {
	return &listInterpreterType{
		opt:         opt,
		tpl:         tpl,
		ipFormatter: f,
	}
}

func (l *listInterpreterType) Parse() ([]byte, error) {
	var lp []string
	if strings.ToLower(l.opt.Src) == "alt" {
		lp = loader.ListCrawlerProxyAlt(l.opt.Max)
	}
	if strings.ToLower(l.opt.Src) == "nova" {
		lp = loader.ListCrawlerProxyNova(l.opt.Max)
	}
	if len(lp) == 0 {
		log.Fatalln("no proxy found")
	}
	for i := 1; i <= l.opt.Max; i++ {
		lbl := fmt.Sprintf("#{PROXY%d}", i)
		l.ipFormatter.SetData(lp[i-1])
		fd, err := l.ipFormatter.Exec()
		if err != nil {
			return nil, err
		}
		l.tpl = bytes.Replace(l.tpl, []byte(lbl), []byte(fd.(string)), -1)
	}

	return l.tpl, nil
}

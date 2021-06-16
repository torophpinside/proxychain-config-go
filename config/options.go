package config

import "flag"

type Options struct {
	SearchType string
	Max        int
	Src        string
}

func GetOptions() Options {
	opt := Options{}
	opt.SearchType = "list"
	flag.IntVar(&opt.Max, "max", 19, "Max of IPs proxies")
	//flag.StringVar(&opt.Src, "source", "", "Source to read the list when search is 'list': alt, free, nova, plus, spysone, all")
	flag.StringVar(&opt.Src, "source", "", "Source to read the list when search is 'list': alt, free, nova, spysone, all")
	flag.Parse()

	return opt
}

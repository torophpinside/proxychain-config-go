package config

import "flag"

type Options struct {
	SearchType string
	Max        int
	Src        string
}

func GetOptions() Options {
	opt := Options{}
	flag.StringVar(&opt.SearchType, "search", "random", "Type of search proxylst: random/list")
	flag.IntVar(&opt.Max, "max", 9, "Max of IPs proxies")
	flag.StringVar(&opt.Src, "source", "", "Source to read the list when search is 'list': nova,alt")
	flag.Parse()

	return opt
}

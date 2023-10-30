package main

import "github.com/Akai66/pandora/internal/design-pattern/proxy"

func main() {
	name := "jack"
	//直接在站点买票
	s := proxy.NewStation(2)
	proxy.Buy(s, name)

	//通过代理买票
	sp := proxy.NewStationProxy(s)
	proxy.Buy(sp, name)
	proxy.Buy(sp, name)
}

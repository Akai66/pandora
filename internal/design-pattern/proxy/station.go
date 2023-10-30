//代理模式：可以为另一个对象提供一个替身或占位符，以控制对这个对象的访问
package proxy

import "fmt"

type Seller interface {
	Sell(string)
}

type Station struct {
	stock int //库存
}

func NewStation(stock int) *Station {
	return &Station{
		stock: stock,
	}
}

func (s *Station) Sell(name string) {
	if s.stock > 0 {
		s.stock--
		fmt.Printf("%s 买了1张票，还剩%d张票\n", name, s.stock)
	} else {
		fmt.Println("抱歉，票已卖完")
	}
}

//代理站点
type StationProxy struct {
	station *Station
}

func NewStationProxy(station *Station) *StationProxy {
	return &StationProxy{
		station: station,
	}
}

func (sp *StationProxy) Sell(name string) {
	fmt.Print("通过代理站点售票:")
	sp.station.Sell(name)
}

func Buy(seller Seller, name string) {
	seller.Sell(name)
}

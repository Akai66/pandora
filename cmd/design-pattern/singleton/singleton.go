package main

import (
	"fmt"
	"github.com/Akai66/pandora/internal/design-pattern/singleton"
)

const pandoraMysqlDns = "root:iam59!z$@tcp(127.0.0.1:3306)/pandora?charset=utf8"

func main() {
	for i := 0; i < 10; i++ {
		go getIns(i, singleton.GetMysqlFunc{
			Name: "Lock",
			Func: singleton.GetMysqlInsOrByLock,
		})
		//go getIns(i, singleton.GetMysqlFunc{
		//	Name: "Once",
		//	Func: singleton.GetMysqlInsOrByOnce,
		//})
	}

}

func getIns(id int, f singleton.GetMysqlFunc) {
	ins, err := f.Func(pandoraMysqlDns)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("name:%s,id:%d,ins:%v\n", f.Name, id, ins)
}

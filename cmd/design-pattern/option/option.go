package main

import (
	"fmt"
	"time"

	"github.com/Akai66/pandora/internal/design-pattern/option"
)

func main() {
	conn, _ := option.NewConnection("127.0.0.1")
	fmt.Printf("%+v\n", conn)

	conn, _ = option.NewConnection("127.0.0.1", option.WithCache(true))
	fmt.Printf("%+v\n", conn)

	conn, _ = option.NewConnection("192.168.1.1", option.WithCache(true), option.WithTimeout(5*time.Second))
	fmt.Printf("%+v\n", conn)
}

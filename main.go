package main

import (
	"github.com/Gosewinckel/Gator/internal/config"
	"fmt"
)

func main() {
	conf := config.Read()
	err := conf.SetUser("lew")
	if err != nil {
		fmt.Println(err)
		return
	}
	conf = config.Read()
	fmt.Println(conf)
}

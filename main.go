package main

import (
	"fmt"
	"goBlog/core"
	"goBlog/global"
)

func main() {
	core.InitConf()
	fmt.Println(global.Config)
}

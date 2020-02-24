package main

import (
	_ "github.com/kaituanwang/sdrms/routers"
	_ "github.com/kaituanwang/sdrms/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

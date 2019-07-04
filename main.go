package main

import (
	"github.com/gogf/gf/g"
	_ "github.com/piaohao/gf-admin/boot"
	_ "github.com/piaohao/gf-admin/router"
)

func main() {
	g.Server().Run()
}

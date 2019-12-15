package main

import (
	"gocho/dao"
	"gocho/web"
)

func main() {
	dao.InitDb()
	web.RegisterHandlers()
}

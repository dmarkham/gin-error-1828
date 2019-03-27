package router

import (
	"github.com/dmarkham/gin-error-1828/controller/admin"
	"github.com/dmarkham/gin-error-1828/controller/home"
)

type HandlerFunc func(*string)

func RouterInit() {
	homeIndex := &home.IndexController{}
	GET(homeIndex.Index)
	adminIndex := &admin.IndexController{}
	GET(adminIndex.Index)
}
func GET(handlers ...HandlerFunc) {
	return
}

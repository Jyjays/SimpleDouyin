package main

import (
	"github.com/Jyjays/SimpleDouyin/initialize"
	"github.com/Jyjays/SimpleDouyin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	initialize.InitializeMysql()         

	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}

package main

import (
	"github.com/Jyjays/SimpleDouyin/initialize"
	"github.com/gin-gonic/gin"
)

func main() {
	initialize.InitializeMysql()

	r := gin.Default()

	r.Run()
}

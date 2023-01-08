package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wachayathorn/golang-basic-structure/pkg/handler"
	"github.com/wachayathorn/golang-basic-structure/pkg/service"
)

func main() {
	r := gin.Default()

	svc := service.InitService()
	handler.InitHandler(r, svc)

	r.Run()
}

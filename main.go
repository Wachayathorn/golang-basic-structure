package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/wachayathorn/golang-basic-structure/pkg/config"
	"github.com/wachayathorn/golang-basic-structure/pkg/handler"
	"github.com/wachayathorn/golang-basic-structure/pkg/service"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "local", "Environment")
	flag.Parse()
	c := config.Config{}
	port, err := c.Load(env)
	if err != nil {
		log.Fatalf("Load config error:%s", err.Error())
	}

	log.SetFormatter(&log.TextFormatter{})

	r := gin.Default()
	svc := service.InitService()
	handler.InitHandler(r, svc)

	log.Infof("Server start port:%s\n", port)
	r.Run()
}

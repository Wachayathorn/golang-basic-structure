package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/wachayathorn/golang-basic-structure/pkg/config"
	"github.com/wachayathorn/golang-basic-structure/pkg/db/postgresql"
	"github.com/wachayathorn/golang-basic-structure/pkg/handler"
	"github.com/wachayathorn/golang-basic-structure/pkg/repository"
	"github.com/wachayathorn/golang-basic-structure/pkg/service"
)

// @title API Document
// @version 1.0
// @description		This is a sample server celler server.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes http
func main() {
	r := gin.Default()
	r.GET("/", HealthCheck)

	// Init configuration
	var env string
	flag.StringVar(&env, "env", "local", "Environment")
	flag.Parse()
	c := config.Config{}
	port, err := c.Load(env)
	if err != nil {
		log.Fatalf("Load config error:%s", err.Error())
	}

	fmt.Printf("Config : %+v", c)

	// Init logrus
	log.SetFormatter(&log.TextFormatter{})

	// Init database
	postgres, err := postgresql.Init(postgresql.PostgresConnection{
		Host:     c.Postgresql.Host,
		Port:     c.Postgresql.Port,
		User:     c.Postgresql.Username,
		Password: c.Postgresql.Password,
		DBname:   c.Postgresql.DDname,
		Timezone: c.Postgresql.Timezone,
	})
	if err != nil {
		log.Fatalf("Init postgresql got error:%s", err.Error())
	}

	// Init route
	repo := repository.Init(postgres)
	svc := service.Init(repo)
	handler.Init(r, svc)

	log.Infof("Server start port:%s\n", port)
	r.Run()
}

func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}
	c.JSON(http.StatusOK, res)
}

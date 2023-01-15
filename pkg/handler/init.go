package handler

import (
	"github.com/gin-gonic/gin"
	bookhandler "github.com/wachayathorn/golang-basic-structure/pkg/handler/book"
	"github.com/wachayathorn/golang-basic-structure/pkg/service"
)

type Handler struct {
	svc *service.Service
}

func Init(r *gin.Engine, svc *service.Service) {
	h := Handler{}
	h.svc = svc
	bookhandler.InitV1(r, svc)
}

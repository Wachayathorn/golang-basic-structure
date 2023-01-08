package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wachayathorn/golang-basic-structure/pkg/handler/book"
	"github.com/wachayathorn/golang-basic-structure/pkg/service"
)

type Handler struct {
	svc *service.Service
}

func InitHandler(r *gin.Engine, svc *service.Service) {
	h := Handler{}
	h.svc = svc
	book.InitApiV1(r, svc)
}

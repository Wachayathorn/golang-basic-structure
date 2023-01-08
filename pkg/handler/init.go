package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wachayathorn/golang-basic-structure/pkg/handler/book"
	"github.com/wachayathorn/golang-basic-structure/pkg/service"
)

type Handler struct {
	svc *service.Service
}

func Init(r *gin.Engine, svc *service.Service) *Handler {
	book.InitV1(r, svc)
	return &Handler{
		svc: svc,
	}
}

package web

import (
	"Demo/internal/service"
	"log"
)

type Handler struct {
	svc *service.Service
}

func NewWebHandler(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) Quotation() {
	response, err := h.svc.SendRequest()
	if err != nil {
		log.Fatalf("报价接口调用失败：%v", err)
		return
	}
	log.Println("报价接口返回数据：", response)
}

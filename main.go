package main

import (
	"Demo/internal/repository"
	"Demo/internal/service"
	"Demo/internal/web"
)

// 记得 go mod tidy 一下

func main() {
	repo := repository.NewRepository()
	svc := service.NewServe(repo)
	h := web.NewWebHandler(svc)
	h.Quotation()
}

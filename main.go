package main

import (
	"Demo/internal/repository"
	"Demo/internal/service"
	"Demo/internal/web"
)

func main() {
	repo := repository.NewRepository()
	svc := service.NewServe(repo)
	h := web.NewWebHandler(svc)
	h.Quotation()
}

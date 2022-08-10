package main

import (
	"github.com/zhangweiii/rank/repositories"
	"github.com/zhangweiii/rank/service"
)

func main() {
	repo := repositories.New(repositories.WithSharedSize(1000))
	service := service.NewService(repo)
	_ = service
}

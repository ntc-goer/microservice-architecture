package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ntc-goer/microservice-examples/inmem-discovery/config"
	"github.com/ntc-goer/microservice-examples/inmem-discovery/handler"
	"github.com/ntc-goer/microservice-examples/inmem-discovery/store"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Load config file fail")
	}
	go store.StartCheck()
	h := handler.NewHandler()
	r := gin.Default()
	r.POST("/register", h.RegisterService)
	r.POST("/update-health", h.UpdateHealth)
	r.GET("/services", h.GetServices)
	r.GET("/discover", h.Discover)
	r.Run(fmt.Sprintf("%s:%s", cfg.Address, cfg.Port))
}

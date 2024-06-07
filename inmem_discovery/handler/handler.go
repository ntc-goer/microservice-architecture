package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ntc-goer/microservice-examples/inmem-discovery/store"
	"net/http"
	"sync"
	"time"
)

type Handler struct {
	sync.Mutex
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterService(ctx *gin.Context) {
	h.Lock()
	defer h.Unlock()
	var req store.InstanceInfo
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("Binding request error %v", err)})
		return
	}
	val, ok := store.InstanceData[req.ID]
	if ok {
		if val.Status == store.OFFLINE {
			val.Status = store.ONLINE
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("Instance is working")})
			return
		}
	} else {
		req.LastHealthCheck = time.Now().Add(time.Duration(req.TTL) * time.Second)
		req.Status = store.ONLINE
		store.InstanceData[req.ID] = &req
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "RegisterService Done",
	})
}

func (h *Handler) UpdateHealth(ctx *gin.Context) {
	h.Lock()
	defer h.Unlock()
	var req store.InstanceInfo
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("Binding request error %v", err)})
		return
	}
	val, ok := store.InstanceData[req.ID]
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("Not found service of instanceID %v", req.ID)})
		return
	}
	val.LastHealthCheck = time.Now().Add(time.Duration(val.TTL) * time.Second)
	val.Status = store.ONLINE
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update health Done",
	})
}

func (h *Handler) GetServices(ctx *gin.Context) {
	h.Lock()
	defer h.Unlock()
	ctx.JSON(http.StatusOK, store.InstanceData)
}

func (h *Handler) Discover(ctx *gin.Context) {
	h.Lock()
	defer h.Unlock()
	service := ctx.Query("service")
	if service == "" {
		ctx.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("invalid service name")})
		return
	}
	serviceLst := make([]*store.InstanceInfo, 0)
	for _, info := range store.InstanceData {
		if info.Service == service && info.Status == store.ONLINE {
			serviceLst = append(serviceLst, info)
		}
	}
	ctx.JSON(http.StatusOK, serviceLst)
}

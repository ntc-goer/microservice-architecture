package store

import (
	"log"
	"time"
)

type InstanceStatus string

const (
	ONLINE  InstanceStatus = "online"
	OFFLINE InstanceStatus = "offline"
)

type InstanceInfo struct {
	ID              string         `json:"id"`
	Service         string         `json:"service"`
	Host            string         `json:"host"`
	Port            string         `json:"port"`
	TTL             int            `json:"ttl"`
	DeregisterAfter int            `json:"deregisterAfter"`
	Status          InstanceStatus `json:"status"`
	LastHealthCheck time.Time      `json:"lastHealthCheck"`
}

var InstanceData = make(map[string]*InstanceInfo)

func StartCheck() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			for s, info := range InstanceData {
				if info.LastHealthCheck.Add(time.Duration(info.DeregisterAfter) * time.Second).Before(time.Now()) {
					delete(InstanceData, s)
					log.Printf("deregister instance %s", s)
				} else if info.LastHealthCheck.Add(time.Duration(info.TTL) * time.Second).Before(time.Now()) {
					info.Status = OFFLINE
					log.Printf("ttl out instance %s", s)
				} else {
					log.Printf("Instance %s is working", s)
				}
			}
		}
	}
}

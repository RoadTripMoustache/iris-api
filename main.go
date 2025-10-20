package main

import (
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/apirouter"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/config"
)

func main() {
	config.LoadConfig()

	go func() {
		apirouter.NewMetricsRouter().Serve()
	}()
	apirouter.New().Serve()
}

package main

import (
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter"
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"github.com/RoadTripMoustache/iris_api/pkg/cron"
)

func main() {
	config.LoadConfig()

	// Start background cron jobs
	cron.Start()

	go func() {
		apirouter.NewMetricsRouter().Serve()
	}()
	apirouter.New().Serve()
}

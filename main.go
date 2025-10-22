package main

import (
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter"
	"github.com/RoadTripMoustache/iris_api/pkg/config"
)

func main() {
	config.LoadConfig()

	go func() {
		apirouter.NewMetricsRouter().Serve()
	}()
	apirouter.New().Serve()
}

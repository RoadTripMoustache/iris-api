// Package cron contains all the methods to manage crons.
package cron

import (
	"context"
	"time"

	imgsvc "github.com/RoadTripMoustache/iris_api/pkg/services/images"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
)

// Manager controls background cron-like jobs.
// It currently runs the orphan images cleanup every hour.
var stopChan chan struct{}

// Start launches the cron manager. It triggers a cleanup immediately and then every hour.
func Start() {
	if stopChan != nil {
		return
	}
	stopChan = make(chan struct{})

	// Run once at startup
	go func() {
		logging.Info("running initial orphan images cleanup", map[string]interface{}{"cron": "images_cleanup"})
		imgsvc.CleanupOrphanImages()
	}()

	// Schedule every hour
	go func() {
		ticker := time.NewTicker(time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				logging.Info("running scheduled orphan images cleanup", map[string]interface{}{"cron": "images_cleanup"})
				imgsvc.CleanupOrphanImages()
			case <-stopChan:
				return
			}
		}
	}()
}

// Stop attempts to gracefully stop the cron manager.
func Stop(ctx context.Context) {
	if stopChan == nil {
		return
	}
	select {
	case stopChan <- struct{}{}:
		// stopped
	case <-ctx.Done():
		return
	}
}

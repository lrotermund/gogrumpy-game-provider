package api

import (
	"log"
	"net/http"
	"time"

	"github.com/Halfi/healthcheck"
)

func initHealthCheck() {
	health := healthcheck.NewHandler()

	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100000))
	health.AddReadinessCheck("readiness", healthcheck.Timeout(func() error {
		return nil
	}, 1*time.Second))

	go func() {
		err := http.ListenAndServe("0.0.0.0:8086", health)
		if err != nil {
			log.Println(err)
		}
	}()
}

package main

import (
	"time"

	"github.com/go-co-op/gocron"
)

func schedulerSeconds(seconds int, functionJob func()) {

	job := gocron.NewScheduler(time.UTC)
	job.Every(seconds).Seconds().Do(functionJob)
	job.StartAsync()
}

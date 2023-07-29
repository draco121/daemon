package main

import (
	"shak-daemon/schedulers"
	"time"
)

func main() {
	scheduler := schedulers.NewScheduler()
	scheduler.ScheduleDiagnostics()
	scheduler.StartSchedulers()
	time.Sleep(time.Minute * 30)
}

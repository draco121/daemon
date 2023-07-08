package main

import "shak-daemon/services"

func main() {
	// scheduler := schedulers.NewScheduler()
	// scheduler.ScheduleDiagnostics()
	// scheduler.StartSchedulers()
	// time.Sleep(time.Minute * 20)
	diagnostics := services.NewDiagnosticsService()
	diagnostics.Process()
}

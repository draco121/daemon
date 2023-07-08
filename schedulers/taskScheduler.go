package schedulers

import (
	"shak-daemon/services"

	"github.com/robfig/cron"
)

type Scheduler struct {
	c *cron.Cron
}

func NewScheduler() Scheduler {
	cron := cron.New()
	scheduler := Scheduler{
		c: cron,
	}
	return scheduler
}

func (s *Scheduler) ScheduleDiagnostics() error {
	s.c.Stop()
	diagnostics := services.NewDiagnosticsService()
	return s.c.AddFunc(diagnostics.Spec.CronString, diagnostics.Process)
}

func (s *Scheduler) StartSchedulers() {
	s.c.Start()
}

package schedulers

import (
	"shak-daemon/services"
	"shak-daemon/utils"

	"github.com/robfig/cron/v3"
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
	_, err := s.c.AddFunc(utils.CronString, diagnostics.Process)
	return err
}

func (s *Scheduler) StartSchedulers() {
	s.c.Start()
}

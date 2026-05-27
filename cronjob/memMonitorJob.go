package cronjob

import (
	"runtime"

	"github.com/sayeneBB/s-ui/logger"
)

type MemMonitorJob struct{}

func NewMemMonitorJob() *MemMonitorJob {
	return &MemMonitorJob{}
}

func (s *MemMonitorJob) Run() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	goroutines := runtime.NumGoroutine()
	heapMB := m.Alloc / 1024 / 1024
	sysMB := m.Sys / 1024 / 1024
	gcPct := m.GCCPUFraction * 100

	logger.Infof("[mem] goroutines=%d heap=%dMB sys=%dMB GC=%.1f%% lastGC=%d",
		goroutines, heapMB, sysMB, gcPct, m.NumGC)

	if goroutines > 2000 {
		logger.Warningf("[mem] high goroutine count: %d (threshold: 2000)", goroutines)
	}
	if heapMB > 300 {
		logger.Warningf("[mem] high heap usage: %dMB", heapMB)
	}
}

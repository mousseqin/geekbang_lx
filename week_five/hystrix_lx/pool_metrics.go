package hystrix_lx

import (
	"geekbang_lx/week_five/hystrix_lx/rolling"
	"sync"
)

type poolMetrics struct {
	Mutex *sync.RWMutex
	Updates chan poolMetricsUpdate

	Name  string
	MaxActiveRequests *rolling.Number
	Executed          *rolling.Number
}

type poolMetricsUpdate struct {
	activeCount int
}

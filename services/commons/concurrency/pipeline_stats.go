package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type StageStats struct {
	Stage Stage

	Processed    int
	Successful   int
	Errors       int
	ErrorsByType map[string]int

	StartTime     time.Time
	LastProcessed time.Time

	ProcessingTimeNs int64
}

func NewStageStats(stage Stage) *StageStats {
	return &StageStats{
		Stage:        stage,
		StartTime:    time.Now(),
		ErrorsByType: make(map[string]int),
	}
}

type PipelineStats struct {
	TotalProcessed int
	TotalErrors    int

	StageStats map[Stage]*StageStats

	PipelineStatsLock sync.Mutex
}

func NewPipelineStats() *PipelineStats {
	return &PipelineStats{
		StageStats: make(map[Stage]*StageStats),
	}
}

func (ps *PipelineStats) RegisterStage(name Stage) {
	ps.PipelineStatsLock.Lock()
	defer ps.PipelineStatsLock.Unlock()

	if _, exists := ps.StageStats[name]; !exists {
		ps.StageStats[name] = NewStageStats(name)
	}
}

func (ps *PipelineStats) RecordSuccess(stage Stage, processTimeNs int64) {
	ps.PipelineStatsLock.Lock()
	defer ps.PipelineStatsLock.Unlock()

	stats, ok := ps.StageStats[stage]
	if !ok {
		return
	}

	stats.Successful++
	stats.Processed++
	stats.LastProcessed = time.Now()
	stats.ProcessingTimeNs += processTimeNs

	ps.TotalProcessed++
}

func (ps *PipelineStats) RecordError(stage Stage, err error, processTimeNs int64) {
	ps.PipelineStatsLock.Lock()
	defer ps.PipelineStatsLock.Unlock()

	stats, ok := ps.StageStats[stage]
	if !ok {
		return
	}

	errType := fmt.Sprintf("%T", err)
	stats.ErrorsByType[errType]++

	stats.Errors++
	stats.Processed++
	stats.LastProcessed = time.Now()
	stats.ProcessingTimeNs += processTimeNs

	ps.TotalErrors++
}

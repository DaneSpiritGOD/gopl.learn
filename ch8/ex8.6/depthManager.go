package main

import (
	"fmt"
	"sync/atomic"
)

type depthManager struct {
	maxDepth        int32
	curDepth        int32
	worksOfCurDepth int32
}

func (dm *depthManager) canIncreaseDepth() bool {
	return dm.curDepth < dm.maxDepth && dm.worksOfCurDepth == 0
}

func (dm *depthManager) canLeave() bool {
	return dm.curDepth == dm.maxDepth && dm.worksOfCurDepth == 0
}

func (dm *depthManager) increaseDepth() {
	dm.curDepth++
}

func (dm *depthManager) addWorks() {
	atomic.AddInt32(&dm.worksOfCurDepth, 1)
}

func (dm *depthManager) removeWorks() {
	atomic.AddInt32(&dm.worksOfCurDepth, -1)
}

func (dm *depthManager) String() string {
	return fmt.Sprintf("%d %d %d", atomic.LoadInt32(&dm.maxDepth), atomic.LoadInt32(&dm.curDepth), atomic.LoadInt32(&dm.worksOfCurDepth))
}

package counter

import "sync"

type SloppyCounter struct {
	globalLock             *sync.Mutex
	localLocks             [4]*sync.Mutex
	localCount             [4]int
	threshold, globalCount int
}

func (sc *SloppyCounter) Init(threshold int) *SloppyCounter {
	var slopCounter = new(SloppyCounter)

	slopCounter.threshold = threshold
	slopCounter.globalCount = 0
	slopCounter.globalLock = new(sync.Mutex)

	for i := 0; i < 4; i++ {
		slopCounter.localLocks[i] = new(sync.Mutex)
	}

	return slopCounter
}

func (sc *SloppyCounter) Update(threadId, amt int) {
	sc.localLocks[threadId].Lock()
	defer sc.localLocks[threadId].Unlock()

	sc.localCount[threadId] += amt
	if sc.localCount[threadId] >= sc.threshold {
		sc.globalLock.Lock()
		defer sc.globalLock.Unlock()

		sc.globalCount += sc.localCount[threadId]
		sc.localCount[threadId] = 0
	}
}

func (sc *SloppyCounter) Get() int {
	sc.globalLock.Lock()
	defer sc.globalLock.Unlock()

	return sc.globalCount
}

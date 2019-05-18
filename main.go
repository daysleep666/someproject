package main

import (
	"fmt"
	"sync"
	"time"
)

var INITTIMESTAMP uint64
var index, WORKID uint64
var lastTimeStamp uint64
var mt sync.Mutex

func init() {
	INITTIMESTAMP = uint64(time.Date(2018, 1, 1, 0, 0, 0, 0, time.Now().Location()).UnixNano())
	WORKID = 1
}

func main() {
	var wg sync.WaitGroup
	var count int64 = 100
	for i := int64(0); i < count; i++ {
		wg.Add(int(1))
		go func() {
			for j := int64(0); j < 1; j++ {
				snowFlake()
				// fmt.Println(snowFlake(), lastTimeStamp, index)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func snowFlake() uint64 {
	mt.Lock()
	defer mt.Unlock()

	timeNow := uint64(time.Now().UnixNano())
	if lastTimeStamp == timeNow {
		index = (index + 1) & 111111111111
		if index == 0 {
			for lastTimeStamp == timeNow {
				timeNow = uint64(time.Now().UnixNano())
			}
		}

	} else {
		index = 0
	}
	lastTimeStamp = timeNow
	fmt.Printf("%v\n", ((timeNow-INITTIMESTAMP)<<22 | WORKID<<12 | index))
	return ((timeNow-INITTIMESTAMP)<<22 | WORKID<<12 | index)
}

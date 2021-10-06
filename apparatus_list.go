package main

import (
	"math"
	"sync/atomic"
	"unsafe"
)

type ApparatusList struct {
	numOfApparatus int
	list []Apparatus
}

func (al ApparatusList) getTimeLeft(now int64) int {
	minWait := math.MaxInt32
	for i , _ := range al.list {
		appa := &al.list[i]
		if appa.busy == 0 {
			return 0
		}
		timeLeft := appa.meal.getTimeLeft(now)
		if minWait > timeLeft {
			minWait = timeLeft
		}
	}
	return minWait
}

func newApparatus(numOfApparatus int) ApparatusList {
	var ret ApparatusList
	ret.numOfApparatus = numOfApparatus
	for i := 0; i < numOfApparatus; i++ {
		ret.list = append(ret.list, Apparatus{})
	}
	return ret
}

func (al ApparatusList)useApparatus(cook Cook, meal *Meal, now int64){

	appa := &al.list[0]
	minWait := math.MaxInt32

	//Get the first oven to finish
	for i , _ := range al.list {
		loopAppa := &al.list[i]
		if loopAppa.busy == 0 {
			minWait = 0
			appa = loopAppa
			break
		}
		timeLeft := loopAppa.meal.getTimeLeft(now)
		if minWait > timeLeft {
			minWait = timeLeft
			appa = loopAppa
		}
	}

	appa.mx.Lock()
	atomic.StoreInt32(&appa.busy,1)
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(appa.meal)), unsafe.Pointer(meal))
	meal.prepare(cook)
	atomic.StoreInt32(&appa.busy,0)
	appa.mx.Unlock()
}

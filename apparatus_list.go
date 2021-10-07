package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
)

type ApparatusList struct {
	numOfApparatus int
	list           []*Apparatus
	listMutex      sync.Mutex
}

func (al *ApparatusList) getTimeLeft(now int64) int {
	minWait := math.MaxInt32
	for i, _ := range al.list {
		appa := al.list[i].getValues()
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

func newApparatus(numOfApparatus int) *ApparatusList {
	ret := new(ApparatusList)
	ret.numOfApparatus = numOfApparatus
	for i := 0; i < numOfApparatus; i++ {
		ret.list = append(ret.list, new(Apparatus))
	}
	return ret
}

func (al *ApparatusList) useApparatus(cook *Cook, meal *Meal, now int64) {


	al.listMutex.Lock()
	appa := al.list[0].getValues()
	minWait := math.MaxInt32

	//Get the first oven to finish
	for i, _ := range al.list {
		loopAppa := al.list[i]
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
	if appa.busy == 1 {
		fmt.Println("THE APPARATUS IS BUSY BRUH")
	} //TODO check state of the cook, when the apparatus is busy, the time left goes negative, because he needs to wait for the apparatus to finish
	al.listMutex.Unlock()

	appa.useApparatus(cook,meal,now)

}

func (al *ApparatusList) getStatus() string {
	ret := ""
	for i, apparatus := range al.list {
		identification := "Id:" + strconv.Itoa(i)
		if apparatus.busy == 1 {
			identification += " Used by cook id:"
			if apparatus.cook != nil {
				identification += strconv.Itoa(apparatus.cook.id)
			}
		} else {
			identification += " Free"
		}
		ret += makeDiv(identification)
	}
	return ret
}

package main

import (
	"math"
	"time"
)

type OrderList struct {
	ovenList  []*Meal
	stoveList []*Meal
	nilList   []*Meal
	orderList []*Order
}

func (orderList OrderList) getMeal(cook Cook) *Meal {
	now := time.Now().Unix()
	overallMin := math.MaxInt32
	var ret *Meal = nil
	ovenTimeLeft := kitchen.ovens.getTimeLeft(now)
	for _, meal := range orderList.ovenList {
		if meal.complexity <= cook.rank {
			timeLeft := meal.getTimeLeft(now) + ovenTimeLeft
			if overallMin > timeLeft {
				overallMin = timeLeft
				ret = meal
			}
		}
	}
	stoveTimeLeft := kitchen.ovens.getTimeLeft(now)
	for _, meal := range orderList.stoveList {
		if meal.complexity <= cook.rank {
			timeLeft := meal.getTimeLeft(now) + stoveTimeLeft
			if overallMin > timeLeft {
				overallMin = timeLeft
				ret = meal
			}
		}
	}
	for _, meal := range orderList.nilList {
		if meal.complexity <= cook.rank {
			timeLeft := meal.getTimeLeft(now)
			if overallMin > timeLeft {
				overallMin = timeLeft
				ret = meal
			}
		}
	}
	return ret
}

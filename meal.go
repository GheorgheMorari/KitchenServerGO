package main

import (
	"sync/atomic"
	"time"
)

type Meal struct {
	prepared      int32
	busy          int32
	timeRequired  int
	complexity    int
	apparatus     int
	preparingTime int64
	parent        *Order
}

func (m Meal) getTimeLeft(now int64) int {
	//now := time.Now().Unix()
	if m.busy == 1 {
		elapsed := int(now - m.preparingTime)
		return m.timeRequired - elapsed
	}
	elapsed := int(now - m.parent.pickUpTime)
	limit := m.parent.maxWait
	priority := m.parent.priority
	return limit - elapsed - m.timeRequired - priority
}

func (m *Meal) prepare() {
	atomic.StoreInt32(&m.busy, 1)
	atomic.StoreInt64(&m.preparingTime,time.Now().Unix())
	time.Sleep(time.Duration(m.timeRequired) * time.Second)
	atomic.StoreInt32(&m.prepared, 1)
	atomic.StoreInt32(&m.busy, 0)
}
func newMeal(parent *Order, id int) *Meal {
	food := menu[id]
	return &Meal{0, 0, food.preparationTime, food.complexity, apparatus[food.cookingApparatus],0, parent}
}

package main

import (
	"sync/atomic"
	"time"
)

type Cook struct {
	id          int
	rank        int
	proficiency int
	name        string
	catchPhrase string
	atWork      int32
}

func (c Cook) startWorking() {
	c.atWork = 1
	for c.atWork == 1 {
		meal := kitchen.orderList.getMeal(c)
		if meal != nil {
			switch meal.apparatus {
			case 0:
				meal.prepare(c)
			case 1:
				kitchen.ovens.useApparatus(c, meal, time.Now().Unix())
			case 2:
				kitchen.stoves.useApparatus(c, meal, time.Now().Unix())
			}
		}
		delivery := kitchen.orderList.getDelivery()
		if delivery != nil {
			kitchen.kitchenWeb.deliver(delivery)
		}
		if meal == nil && delivery == nil{
			//Sleep for one second when there is nothing to do
			time.Sleep(time.Second)
		}
	}
}

func (c Cook) stopWorking() {
	atomic.StoreInt32(&c.atWork, 0)
}


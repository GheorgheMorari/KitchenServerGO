package main

import (
	"math"
	"sync"
	"time"
)

func removeFromArr(arr *[]*Meal, ptr *Meal) {
	index := -1
	for i, meal := range *arr {
		if meal == ptr {
			index = i
			break
		}
	}
	if index != -1 {
		*arr = append((*arr)[:index], (*arr)[index+1:]...)
	}
}

type OrderList struct {
	deliveryMutex sync.Mutex
	mealMutex     sync.Mutex
	ovenList      []*Meal
	stoveList     []*Meal
	nilList       []*Meal
	orderArr      []*Order
}

func (orderList *OrderList) getArray(meal *Meal) *[]*Meal {
	apparatusId := meal.apparatus
	switch apparatusId {
	case 0:
		return &orderList.nilList
	case 1:
		return &orderList.ovenList
	case 2:
		return &orderList.stoveList
	}
	return nil
}

func (orderList *OrderList) addOrder(order *Order) {
	orderList.orderArr = append(orderList.orderArr, order)
	for _, meal := range order.mealList {
		arr := *orderList.getArray(meal)
		arr = append(arr, meal)
	}
}

func (orderList *OrderList) getDelivery() *Delivery {
	//Prevent threads from getting the same delivery
	orderList.deliveryMutex.Lock()
	defer orderList.deliveryMutex.Unlock()

	for i, order := range orderList.orderArr {
		if order.isReady() {
			for _, meal := range order.mealList {
				arr := orderList.getArray(meal)
				removeFromArr(arr, meal)
			}
			orderList.orderArr = append(orderList.orderArr[:i], orderList.orderArr[i+1:]...)
			return newDelivery(order)
		}
	}
	return nil
}

func (orderList *OrderList) getMeal(cook Cook) *Meal {
	//Prevent threads from taking the same meal
	orderList.mealMutex.Lock()
	defer orderList.mealMutex.Unlock()

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

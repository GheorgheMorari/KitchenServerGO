package main

type Order struct {
	mealCounter int
	priority    int
	pickUpTime  int64
	maxWait     int
	mealList    []*Meal
}

func parseOrder(order PostOrder) Order {
	var ret Order
	ret.mealCounter = 0
	ret.priority = order.Priority
	ret.pickUpTime = order.PickUpTime
	ret.maxWait = order.MaxWait
	for _, id := range order.Items {
		ret.mealCounter += 1
		meal := newMeal(&ret,id)
		ret.mealList = append(ret.mealList, meal)
	}
	return ret
}

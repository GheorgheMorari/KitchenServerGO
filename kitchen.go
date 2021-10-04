package main

type Kitchen struct {
	kitchenWeb KitchenWeb
	orderList  OrderList
	ovens      ApparatusList
	stoves     ApparatusList
	cookList   CookList
}

func (k *Kitchen) start() {
	k.ovens = newApparatus(ovenN)
	k.stoves = newApparatus(stoveN)
	k.kitchenWeb.start()
	k.cookList.start()
}

func (k *Kitchen) deliver(delivery *Delivery) {
	k.kitchenWeb.deliver(delivery)
}

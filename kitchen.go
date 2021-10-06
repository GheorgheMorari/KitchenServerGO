package main

import "time"

type Kitchen struct {
	kitchenWeb KitchenWeb
	orderList  OrderList
	ovens      ApparatusList
	stoves     ApparatusList
	cookList   CookList
	connected bool
}

func (k *Kitchen) start() {
	go k.tryConnectDiningHall()
	k.kitchenWeb.start()
}

func (k *Kitchen) tryConnectDiningHall() {
	k.connected = false
	for k.connected {
		if k.kitchenWeb.establishConnection() {
			k.connectionSuccessful()
			break
		} else {
			time.Sleep(time.Second)
		}
	}
}

func (k *Kitchen) deliver(delivery *Delivery) {
	k.kitchenWeb.deliver(delivery)
}

func (k *Kitchen) connectionSuccessful() {
	if k.connected {return}
	k.connected = true
	k.ovens = newApparatus(ovenN)
	k.stoves = newApparatus(stoveN)
	k.cookList.start()
}

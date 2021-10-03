package main


type Kitchen struct {
	kitchenWeb KitchenWeb
	orderList OrderList
	ovens ApparatusList
	stoves ApparatusList
}

func (k *Kitchen) start()  {
	k.ovens = newApparatus(oven_n)
	k.stoves = newApparatus(stove_n)
	k.kitchenWeb.start()
}



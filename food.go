package main

var menu = []Food{pizza,salad,zeama,sswmlc,idwmm,waffles,aubergine,lasagna,burger,gyros}
var apparatus = map[string]int{"":0,"oven":1,"stove":2}
type Food struct {
	id               int
	name             string
	preparationTime  int
	complexity       int
	cookingApparatus string
}

var pizza = Food{1,"pizza",20,2,"oven"}
var salad = Food{2,"salad",10,1, ""}
var zeama = Food{3,"zeama",7,1, "stove"}
var sswmlc = Food{4,"Scallop Sashimi with Meyer Lemon Confit",32,3, ""}
var idwmm = Food{5,"Island Duck with Mulberry Mustard",35,3, "oven"}
var waffles = Food{6,"Waffles",10,1, "stove"}
var aubergine = Food{7,"Aubergine",20,2, ""}
var lasagna = Food{8,"Lasagna",30,2, "oven"}
var burger = Food{9,"Burger",15,1, "oven"}
var gyros = Food{10,"Gyros",15,1, ""}

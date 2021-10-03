package main

import "os"

var diningHallHost = "http://localhost"
const kitchenServerHost = "http://localhost"


const diningHallPort = ":7500"
const kitchenServerPort = ":8000"

const cook_n = 3
const oven_n = 2
const stove_n = 2

var kitchen Kitchen
func main() {
	if args := os.Args; len(args) > 1{
		//Set the docker internal host
		diningHallHost = args[1]
	}
	kitchen.start()
}

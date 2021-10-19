package main

import (
	"os"
	"time"
)

var diningHallHost = "http://localhost"

const diningHallPort = ":7500"
const kitchenServerPort = ":8000"
const cookN = 4
const ovenN = 2
const stoveN = 1


const timeUnit = time.Second

var kitchen Kitchen
func main() {
	if args := os.Args; len(args) > 1{
		//Set the docker internal host
		diningHallHost = args[1]
	}
	kitchen.start()
}

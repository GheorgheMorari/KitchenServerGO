package main

import "math/rand"

type CookList struct {
	cookList      []Cook
	cookIdCounter int
}

func (c CookList) start() {
	c.cookIdCounter = 0
	for i := 0; i < cookN; i++ {
		randomCook := cookPersonas[rand.Intn(len(cookPersonas))]
		randomCook.id = c.cookIdCounter
		c.cookIdCounter++
		if i == 0 {
			randomCook.rank = 3
		}
		c.cookList = append(c.cookList, randomCook)
	}

	for _, cook := range c.cookList {
		go cook.startWorking()
	}
}

var cookPersonas = []Cook{{
	rank:        1,
	proficiency: 1,
	name:        "Jimmy Cook",
	catchPhrase: "YES!",
}, {
	rank:        2,
	proficiency: 2,
	name:        "Andy",
	catchPhrase: "Why am i here?",
}, {
	rank:        1,
	proficiency: 3,
	name:        "Karen",
	catchPhrase: "Abolish the patriarchy",
}, {
	rank:        3,
	proficiency: 2,
	name:        "Vanessa",
	catchPhrase: "The cake is a lie",
}, {
	rank:        3,
	proficiency: 3,
	name:        "Gordon Ramsay",
	catchPhrase: "WHERE IS THE LAMB SAUCE?",
}}

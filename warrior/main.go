package main

import (
	"fmt"
	"math/rand"
)

type warior struct {
	name   string
	health int
	damage int
}

func (w *warior) makeDamage(warior *warior) {
	randomNumber := rand.Intn(11)
	damage := w.damage + randomNumber
	warior.health -= damage
	fmt.Println(w.name, "deal", damage, "damage to ", warior.name)
	fmt.Println(*w, *warior)
}

func run() {
	warrior1 := warior{name: "warrior1", health: 100, damage: 10}
	warrior2 := warior{name: "warrior2", health: 100, damage: 10}
	for {
		(&warrior1).makeDamage(&warrior2)
		if warrior2.health <= 0 {
			fmt.Println("warrior 1 win")
			break
		}
		(&warrior2).makeDamage(&warrior1)
		if warrior1.health <= 0 {
			fmt.Println("warrior 2 win")
			break
		}
	}

	fmt.Println(warrior1, warrior2)
}

func main() {
	run()
}

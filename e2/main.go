package main

import (
	"fmt"
	"math/rand"
	"time"

	"diconium.com/madeifra/go-workshop-3/pokemon"
)

var storage = map[int]pokemon.Pokemon{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		go func(id int) {
			if p, ok := fromPokeballs(id); ok {
				fmt.Println("In Pokeballs")
				fmt.Println(p)
			}
		}(id)
		go func(id int) {
			if p, ok := catch(id); ok {
				fmt.Println("Catch Pokemon")
				fmt.Println(p)
			}
		}(id)
		time.Sleep(150 * time.Millisecond)

		//fmt.Printf("Pokemon not found id: '%v'\n", id)
	}
	time.Sleep(2 * time.Second)
}

func fromPokeballs(id int) (pokemon.Pokemon, bool) {
	p, ok := storage[id]
	return p, ok
}

func catch(id int) (pokemon.Pokemon, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, p := range pokemon.Wild {
		if p.ID == id {
			storage[id] = p
			return p, true
		}
	}

	return pokemon.Pokemon{}, false
}

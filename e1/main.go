package main

import (
	"fmt"
	"math/rand"
	"time"

	"diconium.com/madeifra/go-workshop-3/pokemon"
)

var belt = map[int]pokemon.Pokemon{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 6; i++ {
		id := rnd.Intn(10) + 1
		if p, ok := fromPokeballs(id); ok {
			fmt.Println("In Pokeballs")
			fmt.Println(p)
			continue
		}
		if p, ok := catch(id); ok {
			fmt.Println("Catch Pokemon")
			belt[id] = p
			fmt.Println(p)
			continue
		}
		fmt.Printf("Pokemon not found id: '%v'", id)
	}
}

func fromPokeballs(id int) (pokemon.Pokemon, bool) {
	p, ok := belt[id]
	return p, ok
}

func catch(id int) (pokemon.Pokemon, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, p := range pokemon.Wild {
		if p.ID == id {
			return p, true
		}
	}

	return pokemon.Pokemon{}, false
}

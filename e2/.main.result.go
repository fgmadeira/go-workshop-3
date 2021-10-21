package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"diconium.com/madeifra/go-workshop-3/pokemon"
)

var storage = map[int]pokemon.Pokemon{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 6; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := fromPokeballs(id); ok {
				fmt.Println("In Pokeballs")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := catch(id); ok {
				fmt.Println("Catch Pokemon")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
		time.Sleep(150 * time.Millisecond)

		//fmt.Printf("Pokemon not found id: '%v'\n", id)
	}
	wg.Wait()
}

func fromPokeballs(id int) (pokemon.Pokemon, bool) {
	b, ok := storage[id]
	return b, ok
}

func catch(id int) (pokemon.Pokemon, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, b := range pokemon.Wild {
		if b.ID == id {
			storage[id] = b
			return b, true
		}
	}

	return pokemon.Pokemon{}, false
}

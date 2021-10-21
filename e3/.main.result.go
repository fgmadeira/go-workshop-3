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
	m := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if p, ok := fromPokeballs(id, m); ok {
				fmt.Println("In Pokeballs")
				fmt.Println(p)
			}
			wg.Done()
		}(id, wg, m)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if p, ok := catch(id, m); ok {
				fmt.Println("Catch Pokemon")
				fmt.Println(p)
			}
			wg.Done()
		}(id, wg, m)
		time.Sleep(150 * time.Millisecond)

		//fmt.Printf("Pokemon not found id: '%v'\n", id)
	}
	wg.Wait()
}

func fromPokeballs(id int, m *sync.RWMutex) (pokemon.Pokemon, bool) {
	m.RLock()
	p, ok := storage[id]
	m.RUnlock()
	return p, ok
}

func catch(id int, m *sync.RWMutex) (pokemon.Pokemon, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, p := range pokemon.Wild {
		if p.ID == id {
			m.Lock()
			storage[id] = p
			m.Unlock()
			return p, true
		}
	}

	return pokemon.Pokemon{}, false
}

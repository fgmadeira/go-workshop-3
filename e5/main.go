package main

import (
	"fmt"
	"sync"
	"time"

	"diconium.com/madeifra/go-workshop-3/pokemon"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan pokemon.Pokemon)//10

	wg.Add(2)
	go func(ch chan pokemon.Pokemon, wg *sync.WaitGroup) {
		p := <-ch
		fmt.Printf("%s atacks with a %s atack!\n", p.Name, p.Type)
		time.Sleep(time.Second)
		wg.Done()
	}(ch, wg)
	go func(ch chan pokemon.Pokemon, wg *sync.WaitGroup) {
		p := pokemon.Wild[1]
		fmt.Printf("Go %s!\n", p.Name)
		ch <- p
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

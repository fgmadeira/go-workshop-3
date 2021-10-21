package main

import (
	"fmt"
	"sync"
	"time"

	"diconium.com/madeifra/go-workshop-3/pokemon"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan pokemon.Pokemon) //10
	ac := make(chan string)
	pc := make(chan string)

	wg.Add(2)
	go func(ch <-chan pokemon.Pokemon, wg *sync.WaitGroup, ac chan string) {
		for p := range ch {
			ac <- p.Type
		}
		close(ac)
		wg.Done()
	}(ch, wg, ac)
	go func(ch chan<- pokemon.Pokemon, wg *sync.WaitGroup, pc chan string) {
		for _, p := range pokemon.Wild {
			pc <- p.Name
			ch <- p
			time.Sleep(time.Second)
		}
		close(pc)
		close(ch)
		wg.Done()
	}(ch, wg, pc)

	for {
		if ac == nil && pc == nil {
			break
		}
		select {
		case p, ok := <-ac:
			fmt.Printf("%s attack!\n", p)
			if !ok {
				ac = nil
			}
		case p, ok := <-pc:
			fmt.Printf("Go %s!\n", p)
			if !ok {
				pc = nil
			}
		}
	}

	wg.Wait()
}

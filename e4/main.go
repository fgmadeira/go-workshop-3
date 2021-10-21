package main

import (
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(wg *sync.WaitGroup) {

	}(wg)
	go func(wg *sync.WaitGroup) {

	}(wg)

	wg.Wait()
}

/*
 The goal of this exercice is to use channels to pass information arround go routines using channels.
 Write a first go routine that outputs <Name> atackes with a <Type> atack! when a pokemon is put in the channel.
 Then Write a second go routine that feeds the channel with a pokemon.

 Hint, you probably need to create a channel!

 Extra: you can loop the Wild and try sending all pokemons to the battle!
*/
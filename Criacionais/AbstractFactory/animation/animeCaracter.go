package main

import "fmt"

type animeCaracter struct{}

func (a *animeCaracter) paint() {
	fmt.Println("Desenhando um personagem de anime.")
}

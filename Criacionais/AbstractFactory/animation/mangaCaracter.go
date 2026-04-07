package main

import "fmt"

type mangaCaracter struct{}

func (a *mangaCaracter) paint() {
	fmt.Println("Desenhando um personagem de desenho mangá.")
}

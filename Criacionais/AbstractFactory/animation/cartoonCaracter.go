package main

import "fmt"

type cartoonCaracter struct{}

func (a *cartoonCaracter) paint() {
	fmt.Println("Desenhando um personagem de desenho animado.")
}

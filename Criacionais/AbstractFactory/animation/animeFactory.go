package main

type animeFactory struct{}

func (a *animeFactory) draw() Caracter {
	return &animeCaracter{}
}

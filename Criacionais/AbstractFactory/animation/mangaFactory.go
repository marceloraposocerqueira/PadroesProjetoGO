package main

type mangaFactory struct{}

func (a *mangaFactory) draw() Caracter {
	return &mangaCaracter{}
}

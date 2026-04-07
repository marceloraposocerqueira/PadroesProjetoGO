package main

type cartoonFactory struct{}

func (a *cartoonFactory) draw() Caracter {
	return &cartoonCaracter{}
}

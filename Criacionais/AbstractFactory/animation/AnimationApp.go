package main

import (
	"log"
)

type Application struct {
	factory  animationFactory
	caracter Caracter
}

func NewApplication(f animationFactory) *Application {
	return &Application{
		factory: f,
	}
}

func (a *Application) CreateCaracter() {
	a.caracter = a.factory.draw()
}

func (a *Application) Paint() {
	a.caracter.paint()
}

func main() {
	animationType := "manga"

	log.Printf("Criando um personagem de %s...\n", animationType)

	var factory animationFactory

	switch animationType {
	case "manga":
		factory = &mangaFactory{}
	case "anime":
		factory = &animeFactory{}
	case "cartoon":
		factory = &cartoonFactory{}
	default:
		log.Fatal("Escolha um tipo de animação válido: manga, anime ou cartoon.")
	}

	app := NewApplication(factory)
	app.CreateCaracter()
	app.Paint()
}

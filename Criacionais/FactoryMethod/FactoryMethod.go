package main

import "fmt"

// Produto: Transporte
type Transporte interface {
    Entregar()
}

// Produto Concreto: Carro
type Carro struct{}

func (c *Carro) Entregar() {
    fmt.Println("Entrega realizada por carro.")
}

// Produto Concreto: Bicicleta
type Bicicleta struct{}

func (b *Bicicleta) Entregar() {
    fmt.Println("Entrega realizada por bicicleta.")
}

// Creator: Factory de Transporte
type TransporteFactory interface {
    CriarTransporte() Transporte
}

// Creator Concreto: Factory de Carro
type CarroFactory struct{}

func (f *CarroFactory) CriarTransporte() Transporte {
    return &Carro{}
}

// Creator Concreto: Factory de Bicicleta
type BicicletaFactory struct{}

func (f *BicicletaFactory) CriarTransporte() Transporte {
    return &Bicicleta{}
}

// Função cliente
func realizarEntrega(factory TransporteFactory) {
    transporte := factory.CriarTransporte()
    transporte.Entregar()
}

func main() {
    var factory TransporteFactory

    // Usando a factory de carro
    factory = &CarroFactory{}
    realizarEntrega(factory)

    // Usando a factory de bicicleta
    factory = &BicicletaFactory{}
    realizarEntrega(factory)
}
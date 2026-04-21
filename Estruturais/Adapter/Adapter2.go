package main

import (
	"fmt"
	"math"
)

// --- Interfaces e Estruturas Originais ---

// RoundPegInterface define o comportamento que o RoundHole espera.
type RoundPegInterface interface {
	GetRadius() float64
}

// RoundHole representa o buraco redondo.
type RoundHole struct {
	radius float64
}

func NewRoundHole(radius float64) *RoundHole {
	return &RoundHole{radius: radius}
}

func (r *RoundHole) GetRadius() float64 {
	return r.radius
}

// Fits verifica se um objeto que implementa RoundPegInterface cabe no buraco.
func (r *RoundHole) Fits(peg RoundPegInterface) bool {
	return r.GetRadius() >= peg.GetRadius()
}

// RoundPeg é o pino redondo padrão.
type RoundPeg struct {
	radius float64
}

func NewRoundPeg(radius float64) *RoundPeg {
	return &RoundPeg{radius: radius}
}

func (r *RoundPeg) GetRadius() float64 {
	return r.radius
}

// --- Estrutura Incompatível ---

// SquarePeg é o pino quadrado, que não possui o método GetRadius().
type SquarePeg struct {
	width float64
}

func NewSquarePeg(width float64) *SquarePeg {
	return &SquarePeg{width: width}
}

func (s *SquarePeg) GetWidth() float64 {
	return s.width
}

// --- Adaptador ---

// SquarePegAdapter permite que um SquarePeg seja tratado como um RoundPegInterface.
type SquarePegAdapter struct {
	peg *SquarePeg
}

func NewSquarePegAdapter(peg *SquarePeg) *SquarePegAdapter {
	return &SquarePegAdapter{peg: peg}
}

// GetRadius implementa a interface RoundPegInterface, adaptando a lógica.
func (a *SquarePegAdapter) GetRadius() float64 {
	// O raio de um círculo que engloba um quadrado é metade da diagonal do quadrado.
	return a.peg.GetWidth() * math.Sqrt(2) / 2
}

// --- Execução ---

func main() {
	hole := NewRoundHole(5)
	rpeg := NewRoundPeg(5)
	fmt.Printf("Pino redondo cabe? %v\n", hole.Fits(rpeg)) // true

	smallSqPeg := NewSquarePeg(5)
	largeSqPeg := NewSquarePeg(10)

	// Criando os adaptadores
	smallSqAdapter := NewSquarePegAdapter(smallSqPeg)
	largeSqAdapter := NewSquarePegAdapter(largeSqPeg)

	fmt.Printf("Pino quadrado pequeno cabe? %v\n", hole.Fits(smallSqAdapter)) // true
	fmt.Printf("Pino quadrado grande cabe? %v\n", hole.Fits(largeSqAdapter))  // false
}

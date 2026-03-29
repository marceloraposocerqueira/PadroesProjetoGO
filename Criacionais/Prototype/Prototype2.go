package main

import "fmt"

// Shape funciona como a nossa "interface" de Protótipo.
type Shape interface {
	Clone() Shape
	GetInfo() string
}

// BaseShape contém os campos comuns.
// Em Go, usamos composição (embedding) para simular a herança de campos.
type BaseShape struct {
	X     int
	Y     int
	Color string
}

// Rectangle é um protótipo concreto.
type Rectangle struct {
	BaseShape // Embedding
	Width     int
	Height    int
}

// Clone para Rectangle cria uma nova instância e copia os valores.
func (r *Rectangle) Clone() Shape {
	return &Rectangle{
		BaseShape: BaseShape{
			X:     r.X,
			Y:     r.Y,
			Color: r.Color,
		},
		Width:  r.Width,
		Height: r.Height,
	}
}

func (r *Rectangle) GetInfo() string {
	return fmt.Sprintf("Retângulo [X:%d, Y:%d, Cor:%s, L:%d, A:%d]", r.X, r.Y, r.Color, r.Width, r.Height)
}

// Circle é outro protótipo concreto.
type Circle struct {
	BaseShape
	Radius int
}

func (c *Circle) Clone() Shape {
	return &Circle{
		BaseShape: BaseShape{
			X:     c.X,
			Y:     c.Y,
			Color: c.Color,
		},
		Radius: c.Radius,
	}
}

func (c *Circle) GetInfo() string {
	return fmt.Sprintf("Círculo [X:%d, Y:%d, Cor:%s, Raio:%d]", c.X, c.Y, c.Color, c.Radius)
}

// --- Código Cliente ---

func main() {
	var shapes []Shape

	// Criando um círculo inicial
	circle := &Circle{
		BaseShape: BaseShape{X: 10, Y: 10, Color: "Vermelho"},
		Radius:    20,
	}
	shapes = append(shapes, circle)

	// Clonando o círculo
	anotherCircle := circle.Clone()
	anotherCircle.(*Circle).Radius = 30
	shapes = append(shapes, anotherCircle)

	// Criando um retângulo
	rect := &Rectangle{
		BaseShape: BaseShape{X: 5, Y: 5, Color: "Azul"},
		Width:     10,
		Height:    20,
	}
	shapes = append(shapes, rect)

	// Executando a lógica de negócio (clonagem polimórfica)
	businessLogic(shapes)
}

func businessLogic(shapes []Shape) {
	var shapesCopy []Shape

	fmt.Println("Clonando a lista de formas:")
	for _, s := range shapes {
		cloned := s.Clone()
		shapesCopy = append(shapesCopy, cloned)
		fmt.Printf("Original: %s | Clone: %p (endereço diferente)\n", s.GetInfo(), cloned)
	}
}

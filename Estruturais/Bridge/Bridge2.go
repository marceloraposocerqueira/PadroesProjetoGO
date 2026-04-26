package main

import "fmt"

// 1. Interface da Implementação
type Device interface {
	IsEnabled() bool
	Enable()
	Disable()
	SetVolume(percent int)
}

// 2. Implementações Concretas
type Tv struct{}

func (t *Tv) IsEnabled() bool       { return true }
func (t *Tv) Enable()               { fmt.Println("TV ligada") }
func (t *Tv) Disable()              { fmt.Println("TV desligada") }
func (t *Tv) SetVolume(percent int) { fmt.Printf("Volume TV: %d\n", percent) }

type Radio struct{}

func (r *Radio) IsEnabled() bool       { return false }
func (r *Radio) Enable()               { fmt.Println("Rádio ligado") }
func (r *Radio) Disable()              { fmt.Println("Rádio desligado") }
func (r *Radio) SetVolume(percent int) { fmt.Printf("Volume Rádio: %d\n", percent) }

// 3. Abstração
type Remote struct {
	device Device // A PONTE
}

func (r *Remote) TogglePower() {
	if r.device.IsEnabled() {
		r.device.Disable()
	} else {
		r.device.Enable()
	}
}

// 4. Abstração Refinada
// Usamos composição para simular o comportamento de herança
type AdvancedRemote struct {
	remote Remote
}

func (ar *AdvancedRemote) Mute() {
	ar.remote.device.SetVolume(0)
}

func main() {
	// Uso com TV
	tv := &Tv{}
	basicRemote := &Remote{device: tv}
	basicRemote.TogglePower()

	// Uso com Rádio
	radio := &Radio{}
	advRemote := &AdvancedRemote{remote: Remote{device: radio}}
	advRemote.Mute()
}

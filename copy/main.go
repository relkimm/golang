package main

import "fmt"

type Pos interface {
	Print()
	Copy() Pos
	Move(x, y, z int)
	Rename(name string)
}

type pos struct {
	X, Y, Z int
	Name string
}

func (p *pos) Print() {
	fmt.Printf("%s (%d, %d, %d)\n", p.Name, p.X, p.Y, p.Z)
}

func (p *pos) Copy() Pos {
	copied := NewPos(p.Name, p.X, p.Y, p.Z)
	return copied
}

func (p *pos) Move(x, y, z int) {
	p.X = x
	p.Y = y
	p.Z = z
}

func (p *pos) Rename(name string) {
	p.Name = name
}

func NewPos(name string, x, y, z int) Pos {
	p := new(pos)
	p.Name = name
	p.X = x
	p.Y = y
	p.Z = z
	return p
}

func main() {
	p1 := NewPos("Donggu", 0, 0, 0)
	p2 := p1
	p2.Move(1, 1, 1)
	p1.Print()
	p2.Print()

	p3 := p1.Copy()
	p3.Rename("Segeun")
	p3.Move(2, 2, 2)
	p1.Print()
	p3.Print()
}
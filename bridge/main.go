package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Name() string
}

type circle struct {
	radius float64
}

type rectangle struct {
	w, h float64
}

func NewCircle(r float64) Shape {
	c := new(circle)
	c.radius = r
	return c
}

func (c *circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c *circle) Name() string {
	return "circle"
}

func NewRectangle(w, h float64) Shape {
	r := new(rectangle)
	r.w = w
	r.h = h
	return r
}

func (r *rectangle) Area() float64 {
	return r.w * r.h
}

func (r *rectangle) Name() string {
	return "rectangle"
}

type DrawerImpl interface {
	paint(Shape)
}

type vector struct{}
type raster struct{}

func NewVector() DrawerImpl {
	return new(vector)
}

func (v *vector) paint(shape Shape) {
	fmt.Printf("Draw %s mathmatically taking area of %f\n", shape.Name(), shape.Area())
}

func NewRaster() DrawerImpl {
	return new(raster)
}

func (r *raster) paint(shape Shape) {
	fmt.Printf("Draw %s dot by dot taking area of %f\n", shape.Name(), shape.Area())
}

type Drawer interface {
	DrawShape()
}

type drawer struct {
	impl  DrawerImpl
	shape Shape
}

func NewDrawer(impl DrawerImpl, shape Shape) Drawer {
	d := new(drawer)
	d.impl = impl
	d.shape = shape
	return d
}

func (d *drawer) DrawShape() {
	d.impl.paint(d.shape)
}

func main() {
	vector := NewVector()
	raster := NewRaster()

	circle := NewCircle(3.0)
	rectangle := NewRectangle(2.5, 4.0)

	drawers := []Drawer{
		NewDrawer(vector, circle),
		NewDrawer(raster, circle),
		NewDrawer(vector, rectangle),
		NewDrawer(raster, rectangle),
	}

	for _, drawer := range drawers {
		drawer.DrawShape()
	}
}

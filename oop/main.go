package main

import "fmt"

type Animal interface {
	Sounds()
}

type Dog struct {
	name string
}

func NewDog(name string) Dog {
	return Dog{name: name}
}

func (d Dog) Sounds() {
	fmt.Printf("%s : wolf! wolf!\n", d.name)
}

type Cat struct { 
	name string
}

func NewCat(name string) *Cat {
	return &Cat{name: name}
}

func (c *Cat) Sounds() {
	fmt.Printf("%s : meow.\n", c.name)
}

func sayHello(a Animal) {
	fmt.Print("Hello!\n")
	a.Sounds()
}

func main() {
	dog := NewDog("Charlie")
	dog2 := dog
	dog2.name = "Happy"
	
	dog.Sounds()
	dog2.Sounds()
	
	cat := NewCat("Barley")
	cat2 := cat
	cat2.name = "Sad"

	cat.Sounds()
	cat2.Sounds()

	var a Animal
	a = NewDog("Duck")
	sayHello(dog)
	sayHello(cat)
	sayHello(a)
}

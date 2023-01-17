package main

import (
	"fmt"
	"sync"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func someFunc() {
	printNum := func(num int) {
		fmt.Printf("It is %d\n", num)
	}
	printNum(3)
}

func concurrencyFunc() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
		wg.Add(1)
		go func() {
			fmt.Printf("Job %d finished.\n", <-ch)
			wg.Done()
		}()
	}
	wg.Wait()
}

type Product struct {
	Identifier int
	Name       string
	User       string
}

func (p Product) Info() string {
	return fmt.Sprintf("%s (ID:%d), owned by %s", p.Name, p.Identifier, p.User)
}

func Factory(name string) func(string) Product {
	p := Product{}
	p.Name = name
	i := 0

	return func(user string) Product {
		p.User = user
		p.Identifier = i
		i++
		return p
	}
}

func main() {
	var f func(int, int) int
	a, b := 10, 5

	f = add
	fmt.Println(f(a, b))
	f = sub
	fmt.Println(f(a, b))

	someFunc()
	concurrencyFunc()

	airpodsFactory := Factory("air-pods")
	buzzFactory := Factory("galaxy-buzz")

	p := airpodsFactory("Park")
	fmt.Println(p.Info())

	l := airpodsFactory("Lee")
	fmt.Println(l.Info())

	k := buzzFactory("Kim")
	fmt.Println(k.Info())

	j := buzzFactory("Jeong")
	fmt.Println(j.Info())

}

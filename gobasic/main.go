package main

import (
	"fmt"
	"gobasic/customer"
	"unicode/utf8"
)

func main() {
	w := "เวิลด์"
	fmt.Printf("Hello %v\n", w)
	x := 42 //var x int = 42
	if x >= 30 {
		fmt.Printf("The answer is %d\n", x)
	}
	var y [3]int = [3]int{1, 2}
	fmt.Printf("%v\n", y)
	fmt.Printf("%#v\n", y)
	fmt.Printf("%d\n", y[2])
	z := [3]int{1, 2}
	fmt.Printf("%v\n", z)
	//slice
	sl := []int{1, 2, 4}
	sl = append(sl, 5)
	fmt.Printf("%v, %v\n", sl, len(sl))
	println(utf8.RuneCountInString(w))
	sll := sl[1:]
	fmt.Printf("%v, %v\n", sll, len(sll))
	//map
	countries := map[string]string{}
	countries["th"] = "thailand"
	println(countries["th"])
	country, ok := countries["jp"]
	if ok {
		println(country)
	} else {
		println("no value from countries map")
	}
	//for
	for i := 0; i < len(sll); i++ {
		println(sll[i])
	}
	i := 0
	for i < len(sll) {
		println(sll[i])
		i++
	}
	for _, v := range sll {
		println(v)
	}
	println(sum(5, 5))
	res, _, _ := sum(5, 5)
	println(res)
	ix := func(a, b int) int {
		return a - b
	}
	println(ix(11, 5))
	cal(func(i1, i2 int) int { return i1 * i2 })
	cal(add)
	println(sum2(1, 2, 2, 5))

	//package
	println(customer.Name)
	println(customer.Hello())

	//pointer
	var a int = 10
	var b *int = &a
	fmt.Println(&a, b)
	fmt.Println(a, *b)
	*b = 20
	fmt.Println(a, *b)

	//struct
	p := Person{"test", 17}
	println(p.Name, p.Age)
	//method
	println(p.Hello())
	//encapsulate
	p2 := Person2{}
	p2.SetName("test2")
	println(p2.GetName())
}

func sum(x, y int) (int, string, bool) {
	return x + y, "DONE", true
}

func sum2(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}

func add(x, y int) int {
	return x + y
}

func cal(f func(int, int) int) {
	res := f(10, 10)
	println(res)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Hello() string {
	return "Hello " + p.Name
}

type Person2 struct {
	name string
}

func (p Person2) GetName() string {
	return p.name
}

func (p *Person2) SetName(name string) {
	p.name = name
}

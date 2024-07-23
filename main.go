package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"reflect"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func nakedReturn() (x, y string) {
	x = "hello"
	y = "world"
	return
}

var (
	Tobe   bool       = true
	MaxInt uint64     = 1<<64 - 1
	Z      complex128 = cmplx.Sqrt(-5 + 12i)
)

type vertex struct {
	x int
	y int
}

type Cal interface {
	Add() int
}

type numbers struct {
	a int
	b int
}

func (n numbers) Add() int {
	return n.a + n.b
}

func (n numbers) String() string {
	return fmt.Sprintf("this is string %v %v", n.a, n.b)
}
func (n numbers) Erorr() string {
	return fmt.Sprintf("values are not int %v %v", int(n.a), int(n.b))
}

func genericAddition[T interface{ int | float32 | string }](a T, b T) T {
	return a + b
}
func main() {

	ch := make(chan int, 10)
	go sumOfFewNumbers(ch)

	/* 	for i:=0; i<10; i++{
		fmt.Println("this is channel value in loop", i)
	} */

	for i := range ch {
		fmt.Println("this is from the channel", i)
	}

	/* fmt.Println("this is the sum from the channel", <-ch) */

	a := rand.Float32()
	b := fmt.Sprintf("hola %f", a)

	fmt.Println("hellow", a, "ola", b, "and", math.Pi)

	c := add(10, 11)
	fmt.Println("c", c)

	e, d := swap("world", "hello")
	fmt.Println(e, d)

	f, g := nakedReturn()
	fmt.Println(f, g)

	fmt.Printf("type %T value %v\n", Tobe, Tobe)
	fmt.Printf("type %T value %v\n", MaxInt, MaxInt)
	fmt.Printf("type %T value %v\n", Z, Z)

	fmt.Printf("bitwise and %v\n", 4&5)

	fmt.Printf("bitwise or %v\n", 4|5)

	fmt.Printf("bitwise xor %v\n", 4^5)

	// pointers in go

	i := 24
	j := &i

	fmt.Printf("%v %v\n", &i, *j)

	// struct
	v := vertex{1, 3}
	fmt.Println(v)

	fmt.Println(&v.x, &v.y)

	//TimeTaken()

	n := numbers{101, 20}
	fmt.Printf("%T", n)
	add := n.Add()
	fmt.Printf("%v", add)
	fmt.Println(n)

	val := genericAddition(11, 22)
	fmt.Println(val)

	val2 := genericAddition[float32](0.1, 0.2)
	fmt.Println(val2, "this is the second value")

	val3 := genericAddition("abba", "dabba")
	fmt.Println(val3)

	ref1 := reflect.ValueOf(val3)
	ref2 := reflect.TypeOf(val3)

	fmt.Println("this is reflect package usage", ref1, ref2)
}

func sumOfFewNumbers(ch chan int) {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
		ch <- sum

	}
	fmt.Printf("this is sum %v ---", sum)
	close(ch)
}

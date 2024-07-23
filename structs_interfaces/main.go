package main

import "fmt"

type Matter struct {
	Mass float64
}

type Force interface {
	calcForce(float64) float64
}

func (m Matter) calcForce(acc float64) float64 {
	return m.Mass * acc
}

func (m Matter) forceWithGravity(gravity float64) float64 {
	return m.Mass * gravity

}
func CreateMatterWithForce(mass float64) Force {
	return Matter{
		Mass: mass,
	}

}
func CreateMatter(mass float64) Matter {
	return Matter{
		Mass: mass,
	}

}

func main() {
	fmt.Println("lets get started")
	utrm := CreateMatter(1000)
	ujrm := CreateMatterWithForce(100)
	forceWithGravity := utrm.forceWithGravity(9.8)
	accUjrm := ujrm.calcForce(100)
	fmt.Println("", accUjrm, forceWithGravity)
}

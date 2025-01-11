package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

type Triangle struct {
	point1 Point
	point2 Point
	point3 Point
}

type Square struct {
	initialPoint Point
	sideLength   float64
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t Triangle) getArea() float64 {
	leftHand := (t.point1.x * t.point2.y) + (t.point2.x * t.point3.y) + (t.point3.x * t.point1.y)
	rightHand := (t.point1.y * t.point2.x) + (t.point2.y * t.point3.x) + (t.point3.y * t.point1.x)
	return math.Abs((leftHand - rightHand) / 2)
}

func (t Triangle) getMiddlePoint() Point {
	return Point{(t.point1.x + t.point2.x + t.point3.x) / 3, (t.point1.y + t.point2.y + t.point3.y) / 3}
}

func (s Square) getMiddlePoint() Point {
	return Point{s.initialPoint.x + (s.sideLength / 2), s.initialPoint.y + (s.sideLength / 2)}
}

func (t Triangle) getDistanceToOrigin() float64 {
	middlePoint := t.getMiddlePoint()
	return math.Sqrt(math.Pow(middlePoint.x-0, 2) + math.Pow(middlePoint.y-0, 2))
}

func (s Square) getDistanceToOrigin() float64 {
	middlePoint := s.getMiddlePoint()
	return math.Sqrt(math.Pow(middlePoint.x-0, 2) + math.Pow(middlePoint.y-0, 2))
}

type getDistanceToOrigin interface {
	getDistanceToOrigin() float64
}

type getArea interface {
	getArea() float64
}

type getMiddlePoint interface {
	getMiddlePoint() Point
}

func main() {
	triangle := Triangle{
		point1: Point{0, 0},
		point2: Point{4, 0},
		point3: Point{2, 4},
	}

	square := Square{
		initialPoint: Point{1, 1},
		sideLength:   2,
	}

	fmt.Println("Triangle Properties:")
	fmt.Println("Area:", triangle.getArea())
	fmt.Println("Middle Point:", triangle.getMiddlePoint())
	fmt.Println("Distance to Origin:", triangle.getDistanceToOrigin())

	fmt.Println("\nSquare Properties:")
	fmt.Println("Area:", square.getArea())
	fmt.Println("Middle Point:", square.getMiddlePoint())
	fmt.Println("Distance to Origin:", square.getDistanceToOrigin())

	var shape getDistanceToOrigin

	shape = triangle
	fmt.Println("\nTriangle Distance to Origin (via Interface):", shape.getDistanceToOrigin())

	shape = square
	fmt.Println("Square Distance to Origin (via Interface):", shape.getDistanceToOrigin())
}

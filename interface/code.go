package main
import "fmt"

type Shape interface {
	area() float32
}

type Rectangle struct {
	length, breadth float32
}

type Triangle struct {
	base, height float32
}

func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

func (t Triangle) area() float32 {
	return 0.5 * t.base * t.height
}

func calcRectangleArea(s Shape) float32 {
	return s.area()
}

func calcTriangleArea(t Shape) float32 {
	return t.area()
}

func main() {
	rect := Rectangle {
		length: 10,
		breadth: 20,
	}

	tri := Triangle{10, 20}

	fmt.Println("Area of Rectangle:", calcRectangleArea(rect))
	fmt.Println("Area of Triangle:", calcTriangleArea(tri))
}
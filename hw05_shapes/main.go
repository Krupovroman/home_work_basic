package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() (float64, error)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() (float64, error) {
	if c.Radius <= 0 {
		return 0, errors.New("radius must be positive")
	}
	return math.Pi * c.Radius * c.Radius, nil
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() (float64, error) {
	if r.Width <= 0 || r.Height <= 0 {
		return 0, errors.New("width and height must be positive")
	}
	return r.Width * r.Height, nil
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() (float64, error) {
	if t.Base <= 0 || t.Height <= 0 {
		return 0, errors.New("base and height must be positive")
	}
	return 0.5 * t.Base * t.Height, nil
}

func calculateArea(s Shape) (float64, error) {
	return s.Area()
}

func main() {
	circle := Circle{Radius: 5}
	rect := Rectangle{Width: 10, Height: 5}
	triangle := Triangle{Base: 8, Height: 6}

	shapes := []Shape{circle, rect, triangle}

	for _, shape := range shapes {
		area, err := calculateArea(shape)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}
		switch s := shape.(type) {
		case Circle:
			fmt.Printf("Круг: радиус %.1f\n", s.Radius)
		case Rectangle:
			fmt.Printf("Прямоугольник: ширина %.1f, высота %.1f\n", s.Width, s.Height)
		case Triangle:
			fmt.Printf("Треугольник: основание %.1f, высота %.1f\n", s.Base, s.Height)
		}
		fmt.Printf("Площадь: %.2f\n", area)
	}
}

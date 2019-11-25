//	import $GOPATH/src/github.com/Mishamba/EPAM_GOLANG_COURSE/FourthTask/figures/types
//	Package contains structures descriptions and constructors, interface definition

package types

import (
	"errors"
	"math"
)

const (
	circlePanic   string = " circle can't has negative radius"
	squarePanic   string = " square can't has negative side"
	overflowError string = " result will be bigger, that we can coutn. sorry..."
)

type Figure interface {
	area() (float64, error)
	perimetr() (float64, error)
}

type Circle struct {
	radius int
}

func CircleConstruct(newRadius int) (Circle, error) {
	if newRadius < 0 {
		return Circle{0}, errors.New(circlePanic)
	}
	return Circle{newRadius}, nil
}

func (c Circle) Area() (float64, error) {
	if float64(c.radius) > math.Sqrt(9223372036854775807.0/math.Pi) {
		return 0, errors.New(overflowError)
	}
	return math.Pi * math.Pow(float64(c.radius), 2), nil
}

func (c Circle) Perimetr() (float64, error) {
	if float64(c.radius) > 9223372036854775807.0/2.0*math.Pi {
		return 0, errors.New(overflowError)
	}
	return 2.0 * math.Pi * float64(c.radius), nil
}

type Square struct {
	side int
}

func SquareConstruct(newSide int) (Square, error) {
	if newSide < 0 {
		return Square{0}, errors.New(squarePanic)
	}
	return Square{newSide}, nil
}

func (s Square) Area() (float64, error) {
	if float64(s.side) > math.Sqrt(9223372036854775807.0) {
		return 0, errors.New(overflowError)
	}
	return math.Pow(float64(s.side), 2), nil
}

func (s Square) Perimetr() (float64, error) {
	if float64(s.side) > 9223372036854775807.0/4.0 {
		return 0, errors.New(overflowError)
	}
	return 4.0 * float64(s.side), nil
}

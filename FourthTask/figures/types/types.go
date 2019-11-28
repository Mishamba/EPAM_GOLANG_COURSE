//	import $GOPATH/src/github.com/Mishamba/EPAM_GOLANG_COURSE/FourthTask/figures/types
//	Package contains structures descriptions and constructors, interface definition

package types

import (
	"errors"
	"math"
)

var (
	ErrCircleNegative = errors.New(" circle can't has negative radius")
	ErrSquareNegative = errors.New(" square can't has negative side")
	overflowError     = errors.New(" result will be bigger, that we can coutn. sorry...")
)

type Figure interface {
	Area() (float64, error)
	Perimetr() (float64, error)
}

type Circle struct {
	radius int
}

func NewCircle(newRadius int) (Figure, error) {
	if newRadius < 0 {
		return nil, ErrCircleNegative
	}
	return Circle{newRadius}, nil
}

func (c Circle) Area() (float64, error) {
	if float64(c.radius) > math.Sqrt(math.MaxFloat64/math.Pi) {
		return 0, overflowError
	}
	return math.Pi * math.Pow(float64(c.radius), 2), nil
}

func (c Circle) Perimetr() (float64, error) {
	if float64(c.radius) > math.MaxFloat64/(2.0*math.Pi) {
		return 0, overflowError
	}
	return 2.0 * math.Pi * float64(c.radius), nil
}

type Square struct {
	side int
}

func NewSquare(newSide int) (Figure, error) {
	if newSide < 0 {
		return nil, ErrSquareNegative
	}
	return Square{newSide}, nil
}

func (s Square) Area() (float64, error) {
	if float64(s.side) > math.Sqrt(math.MaxFloat64) {
		return 0, overflowError
	}
	return math.Pow(float64(s.side), 2), nil
}

func (s Square) Perimetr() (float64, error) {
	if float64(s.side) > math.MaxFloat64/4.0 {
		return 0, overflowError
	}
	return 4.0 * float64(s.side), nil
}

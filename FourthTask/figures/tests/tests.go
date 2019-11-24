//	Tests package. Here i'll write some tests and if smth will go wrong i'll tell main pocket to stop his work
package tests

import (
	"errors"
	"math"

	"github.com/Mishamba/EPAM_GOLANG_COURSE/FourthTask/figures/types"
)

//i don't know, what else to test here

func Errors() error {
	_, err := types.CircleConstruct(-5) //negative input
	if err != nil {
		return err
	}

	_, err = types.SquareConstruct(-5) // -//-
	if err != nil {
		return err
	}

	circle, err := types.CircleConstruct(5) // checking returns value
	if err != nil {
		return err
	}
	area, _ := circle.Area()
	if math.Pi*math.Pow(5.0, 2.0) != area {
		return errors.New("Circles area returns wrong value.")
	}
	perimetr, _ := circle.Perimetr()
	if math.Pi*2.0*5.0 != perimetr {
		return errors.New("Circle perimetr returns wrong value.")
	}

	square, err := types.SquareConstruct(5)
	if err != nil {
		return err
	}
	area, _ = square.Area()
	if math.Pow(5.0, 2.0) != area {
		return errors.New("Square area returns wrong value.")
	}
	perimetr, _ = square.Perimetr()
	if 4.0*5.0 != perimetr {
		return errors.New("Square perimetr returns wrong value.")
	}

	return nil
}

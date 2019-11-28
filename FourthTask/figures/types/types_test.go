//	Tests package. Here i'll write some tests and if smth will go wrong i'll tell main pocket to stop his work
package types

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	_, err := NewCircle(-5)
	assert.NotNil(t, err, "Circle construct tried to create negative circle")

	_, err = NewSquare(-5)
	assert.NotNil(t, err, "Square construct tried to create negative square")

	var value float64

	circle, err := NewCircle(5)
	assert.Nil(t, err, "Circle constructor returns error when it shouldn't")

	area, _ := circle.Area()
	value = math.Pi * math.Pow(5.0, 2.0)
	assert.Equal(t, value, area, "Area for circle returns wrong value")

	perimetr, _ := circle.Perimetr()
	value = math.Pi * 2 * 5.0
	assert.Equal(t, value, perimetr, "Perimetr for square returns wrong value")

	square, err := NewSquare(5)
	assert.Nil(t, err, "Square constructor returns error when it shouldn't")

	area, _ = square.Area()
	value = math.Pow(5.0, 2.0)
	assert.Equal(t, value, area, "Area for square returns wrong value")

	perimetr, _ = square.Perimetr()
	value = 4.0 * 5.0
	assert.Equal(t, value, perimetr, "Perimetr for square returns wrong value")

	//i found out, that int can't create overflow exception, so i deleted this checks in types.go and i woun't use this tests
	/*value = math.Sqrt(math.MaxFloat64/math.Pi) + 1
	if circle, err = NewCircle(int(value)); err != nil { // crashs here. err =  circle can't has negative radius
		panic("test crash")
	}
	area, err = circle.Area()
	assert.Equal(t, area, 0, "Circle area exception - overflow check not works")
	assert.NotNil(t, err, "Circle area exception - overflow check not works")

	value = math.MaxFloat64/(2.0*math.Pi) + 1
	if circle, err = NewCircle(int(value)); err != nil {
		panic("test crash")
	}
	perimetr, err = circle.Perimetr()
	assert.Equal(t, perimetr, 0, "Circle perimert exception - overflow check not works")
	assert.NotNil(t, err, "Circle perimert exception - overflow check not works")

	value = math.Sqrt(math.MaxFloat64) + 1
	if square, err = NewCircle(int(value)); err != nil {
		panic("tests crash")
	}
	area, err = square.Area()
	assert.Equal(t, area, 0, "Square area exeption - overflow check not works")
	assert.NotNil(t, err, "Square area exeption - overflow check not works")

	value = math.MaxFloat64/4.0 + 1
	if square, err = NewCircle(int(value)); err != nil {
		panic("tests crash")
	}
	perimetr, err = square.Area()
	assert.Equal(t, perimetr, 0, "Square perimetr exeption - overflow check not works")
	assert.NotNil(t, err, "Square perimetr exeption - overflow check not works")*/
}

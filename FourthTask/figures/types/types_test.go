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

	area := circle.Area()
	value = math.Pi * math.Pow(5.0, 2.0)
	assert.Equal(t, value, area, "Area for circle returns wrong value")

	perimetr := circle.Perimetr()
	value = math.Pi * 2 * 5.0
	assert.Equal(t, value, perimetr, "Perimetr for square returns wrong value")

	square, err := NewSquare(5)
	assert.Nil(t, err, "Square constructor returns error when it shouldn't")

	area = square.Area()
	value = math.Pow(5.0, 2.0)
	assert.Equal(t, value, area, "Area for square returns wrong value")

	perimetr = square.Perimetr()
	value = 4.0 * 5.0
	assert.Equal(t, value, perimetr, "Perimetr for square returns wrong value")
}

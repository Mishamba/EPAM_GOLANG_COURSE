package structures

type Point struct {
	X, Y int 
}

type Square struct {
	Start Point
	A uint
}

func (my_square Square) End() (x_result, y_result int) {
	return my_square.Start.X + int(my_square.A), my_square.Start.Y - int(my_square.A)
}

func (my_square Square) Perimeter() uint {
    return my_square.A * 4
}

func (my_square Square) Area() uint {
    //return uint(Pow(float64(my_square.a), float64(2))) //как использовать эту функцию?
    return my_square.A * my_square.A
}

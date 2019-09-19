package rectangle

import "math"

//Area to find out the area of a rectangle
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

//Diagonal to find out the diagonal of a rectangle
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}

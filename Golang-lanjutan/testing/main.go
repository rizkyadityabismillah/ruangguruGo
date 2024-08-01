package main

import "math"

type Cube struct{
	side float64
}
func (c Cube) Area() float64{
	return 6* math.Pow(c.side , 2)
}
func (c Cube) Circumference() float64{
	return 12  * c.side
}
func (c Cube) Volume() float64{
	return math.Pow(c.side,3)
}
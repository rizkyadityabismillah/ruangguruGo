package calculator

type Calculator struct{
	Base float64
}
func InitCalculator(base float64)*Calculator {
	return &Calculator{Base: base}
}
func (c *Calculator) Add(number float64){
	c.Base +=  number
}
func (c *Calculator) Subtract(number float64){
	c.Base -= number
}
func (c *Calculator) Multiply(number float64){
	c.Base *= number
}
func (c *Calculator) Divide(number float64) error {
	if number == 0{
		return errors.New("Division by zero is not allowed")
	}
	c.Base /= number
	return nil
}
func (c *Calculator) Result(number float64){
	return c.Base
}
func (c *Calculator) Reset(){
	c.Base = 0
}
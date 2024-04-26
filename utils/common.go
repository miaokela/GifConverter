package utils

type Calculator struct{}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

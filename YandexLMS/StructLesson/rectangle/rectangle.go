// Пакет по работе с пряоугольником
package rectangle

type Rectangle struct {
	Width  float64
	Heigth float64
}

func (a Rectangle) Area() float64 {
	return a.Width * a.Heigth
}
func (s Rectangle) Scale(factor float64) (float64, float64) {
	s.Width *= factor
	s.Heigth *= factor
	return s.Width, s.Heigth
}

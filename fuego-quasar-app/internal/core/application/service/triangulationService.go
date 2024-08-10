package service

import (
	"fuego-quasar-app/internal/core/domain/model"
	"fuego-quasar-app/internal/core/domain/port"
)

type TriangulationService struct{}

func NewTriangulationService() port.TriangulationService {
	return &TriangulationService{}
}
func (t *TriangulationService) GetLocation(p1, p2, p3 model.Point, d1, d2, d3 float64) (model.Point, error) {

	A := 2 * (p2.X - p1.X)
	B := 2 * (p2.Y - p1.Y)
	C := d1*d1 - d2*d2 - p1.X*p1.X + p2.X*p2.X - p1.Y*p1.Y + p2.Y*p2.Y
	D := 2 * (p3.X - p2.X)
	E := 2 * (p3.Y - p2.Y)
	F := d2*d2 - d3*d3 - p2.X*p2.X + p3.X*p3.X - p2.Y*p2.Y + p3.Y*p3.Y

	x := (C*E - F*B) / (E*A - B*D)
	y := (C*D - A*F) / (B*D - A*E)

	return model.Point{X: x, Y: y}, nil
}

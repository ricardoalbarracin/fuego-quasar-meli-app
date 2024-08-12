package service

import (
	"fmt"
	"fuego-quasar-app/internal/core/domain/model"
	"fuego-quasar-app/internal/core/domain/port"
	"math"
)

type TriangulationService struct{}

func NewTriangulationService() port.TriangulationService {
	return &TriangulationService{}
}
func (t *TriangulationService) GetLocation(p1, p2, p3 model.Point, d1, d2, d3 float64) (model.Point, error) {

	// Coeficientes para las ecuaciones
	A := 2*p2.X - 2*p1.X
	B := 2*p2.Y - 2*p1.Y
	C := d1*d1 - d2*d2 - p1.X*p1.X + p2.X*p2.X - p1.Y*p1.Y + p2.Y*p2.Y
	D := 2*p3.X - 2*p2.X
	E := 2*p3.Y - 2*p2.Y
	F := d2*d2 - d3*d3 - p2.X*p2.X + p3.X*p3.X - p2.Y*p2.Y + p3.Y*p3.Y

	// Resolver para x y
	denominator := A*E - B*D
	if math.Abs(denominator) < 1e-10 {
		return model.Point{}, fmt.Errorf("las ecuaciones son paralelas o no tienen solución única")
	}

	x := (C*E - B*F) / denominator
	y := (C*D - A*F) / denominator

	return model.Point{X: x, Y: y}, nil
}

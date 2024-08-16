package service

import (
	"errors"
	"fuego-quasar-app/internal/core/domain/model"
	"fuego-quasar-app/internal/core/domain/port"
	"math"
)

type TriangulationService struct {
	logService port.LogService
}

func NewTriangulationService(logService port.LogService) port.TriangulationService {
	return TriangulationService{logService: logService}
}
func (t TriangulationService) GetLocation(p1, p2, p3 model.Point, d1, d2, d3 float64) (model.Point, error) {

	// Construcción de las ecuaciones
	A := 2 * (p2.X - p1.X)
	B := 2 * (p2.Y - p1.Y)
	C := math.Pow(d1, 2) - math.Pow(d2, 2) - math.Pow(p1.X, 2) + math.Pow(p2.X, 2) - math.Pow(p1.Y, 2) + math.Pow(p2.Y, 2)
	D := 2 * (p3.X - p1.X)
	E := 2 * (p3.Y - p1.Y)
	F := math.Pow(d1, 2) - math.Pow(d3, 2) - math.Pow(p1.X, 2) + math.Pow(p3.X, 2) - math.Pow(p1.Y, 2) + math.Pow(p3.Y, 2)

	// Determinante del sistema
	denominator := A*E - B*D
	if denominator == 0 {
		return model.Point{}, errors.New("no hay solución válida para encontrar su ubicacion")
	}

	// Cálculo de x y y
	x := (C*E - F*B) / denominator
	y := (A*F - C*D) / denominator

	// Verificar si el punto (x, y) está en el tercer círculo
	distanceToP3 := math.Sqrt(math.Pow(x-p3.X, 2) + math.Pow(y-p3.Y, 2))
	valDis := distanceToP3 - d3
	if math.Abs(valDis) > 1e-2 {
		return model.Point{}, errors.New("no hay solución válida para encontrar su ubicacion")
	}

	return model.Point{X: x, Y: y}, nil
}

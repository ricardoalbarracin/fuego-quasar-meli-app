package port

import "fuego-quasar-app/internal/core/domain/model"

type TriangulationService interface {
	GetLocation(p1, p2, p3 model.Point, d1, d2, d3 float64) (model.Point, error)
}

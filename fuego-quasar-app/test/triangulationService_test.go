package service

import (
	"fuego-quasar-app/internal/core/application/service"
	"fuego-quasar-app/internal/core/domain/model"
	infraestructure "fuego-quasar-app/internal/infrastructure/log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangulationService_GetLocation(t *testing.T) {
	service := service.NewTriangulationService(infraestructure.NewLog())

	tests := []struct {
		name       string
		p1, p2, p3 model.Point
		d1, d2, d3 float64
		want       model.Point
		wantErr    bool
	}{
		{
			name:    "Caso de error",
			p1:      model.Point{X: -500, Y: -200},
			p2:      model.Point{X: 100, Y: -100},
			p3:      model.Point{X: 500, Y: 100},
			d1:      100.0,
			d2:      115.5,
			d3:      142.7,
			want:    model.Point{X: 0, Y: 0},
			wantErr: true,
		},
		{
			name:    "Caso de OK",
			p1:      model.Point{X: -500, Y: -200},
			p2:      model.Point{X: 100, Y: -100},
			p3:      model.Point{X: 500, Y: 100},
			d1:      549.74,
			d2:      74.53,
			d3:      495.53,
			want:    model.Point{X: 33.32122087499996, Y: -66.63059174999981},
			wantErr: false,
		},
		{
			name:    "Simple case",
			p1:      model.Point{X: 0, Y: 0},
			p2:      model.Point{X: 1, Y: 0},
			p3:      model.Point{X: 0, Y: 1},
			d1:      1.414213562,
			d2:      1,
			d3:      1,
			want:    model.Point{X: 1, Y: 1},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GetLocation(tt.p1, tt.p2, tt.p3, tt.d1, tt.d2, tt.d3)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.InDelta(t, tt.want.X, got.X, 1e-2) || !assert.InDelta(t, tt.want.Y, got.Y, 1e-2) {
				t.Errorf("GetLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

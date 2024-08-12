package service

import (
	"fuego-quasar-app/internal/core/application/service"
	"fuego-quasar-app/internal/core/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangulationService_GetLocation(t *testing.T) {
	service := service.NewTriangulationService()

	tests := []struct {
		name       string
		p1, p2, p3 model.Point
		d1, d2, d3 float64
		want       model.Point
		wantErr    bool
	}{
		{
			name:    "Simple case",
			p1:      model.Point{X: -500, Y: -200},
			p2:      model.Point{X: 100, Y: -100},
			p3:      model.Point{X: 500, Y: 100},
			d1:      100.0,
			d2:      115.5,
			d3:      142.7,
			want:    model.Point{X: -487.29, Y: -1557.01},
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

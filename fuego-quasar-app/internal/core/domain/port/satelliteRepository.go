package port

import "fuego-quasar-app/internal/core/domain/model"

type SatelliteRepository interface {
	Create(satellite *model.Satellites) error
	FindByName(name string) (*model.Satellites, error)
	Delete(name string) error
	FindAll() ([]*model.Satellites, error)
	FindByNames(names []string) ([]*model.Satellites, error)
	DeleteAll() error
}

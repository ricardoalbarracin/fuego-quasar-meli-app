package port

import "fuego-quasar-app/internal/core/domain/model"

type FuegoQuasarService interface {
	ProcessSplitMessage(satellites model.Satellites) error
	ProcessSaveMessages() (model.Response, error)
	ProcessMessages(satellites []model.Satellites) (model.Response, error)
}

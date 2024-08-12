package port

import "fuego-quasar-app/internal/core/domain/model"

type SecretManagerService interface {
	GetSecret(string) (model.Setting, error)
}

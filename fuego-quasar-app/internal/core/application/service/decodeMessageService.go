package service

import "fuego-quasar-app/internal/core/domain/port"

type DecodeMessageService struct{}

func NewDecodeMessageService() port.DecodeMessageService {
	return &DecodeMessageService{}
}

func (d *DecodeMessageService) GetMessage(messages [][]string) (string, error) {
	return "", nil
}

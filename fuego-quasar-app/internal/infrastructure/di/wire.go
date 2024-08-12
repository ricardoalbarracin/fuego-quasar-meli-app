//go:build wireinject
// +build wireinject

package di

import (
	"fuego-quasar-app/internal/core/application/service"
	"fuego-quasar-app/internal/infrastructure/awsSecret"
	"fuego-quasar-app/internal/infrastructure/mongodb"
	"fuego-quasar-app/internal/infrastructure/repository"
	"fuego-quasar-app/internal/interfaces/handler"

	"github.com/google/wire"
)

func InitializeMyService() handler.LambdaHandler {
	wire.Build(awsSecret.NewAWSSecretManagerService, mongodb.NewMongoClient, repository.NewSatelliteRepositoryMongo, handler.NewLambdaHandler, service.NewTriangulationService, service.NewDecodeMessageService, service.NewFuegoQuasarService)
	return handler.LambdaHandler{}
}

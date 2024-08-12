package mongodb

import (
	"context"
	"fmt"
	"fuego-quasar-app/internal/core/domain/port"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(secretManagerService port.SecretManagerService) *mongo.Client {
	secretName := os.Getenv("CONNECTION_SECRET_NAME")
	setting, err := secretManagerService.GetSecret(secretName)
	if err != nil {
		log.Fatalf("Error al conectar con MongoDB: %v", err)
	}
	log.Printf("setting %s", setting.Value)

	// Crear opciones de cliente MongoDB
	clientOptions := options.Client().ApplyURI(setting.Value)

	// Crear cliente MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Error al conectar con MongoDB: %v", err)
	}

	// Verificar la conexión
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatalf("Error al hacer ping a MongoDB: %v", err)
	}

	fmt.Println("Conexión a MongoDB exitosa")
	return client
}

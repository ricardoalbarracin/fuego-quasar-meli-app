package repository

import (
	"context"
	"fuego-quasar-app/internal/core/domain/model"
	"fuego-quasar-app/internal/core/domain/port"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SatelliteRepositoryMongo struct {
	collection *mongo.Collection
}

// NewSatelliteRepositoryMongo crea una nueva instancia de SatelliteRepositoryMongo.
func NewSatelliteRepositoryMongo(client *mongo.Client) port.SatelliteRepository {
	collection := client.Database("XXXXX").Collection("satellites")
	return &SatelliteRepositoryMongo{collection: collection}
}

// Create inserta un nuevo documento en MongoDB.
func (r *SatelliteRepositoryMongo) Create(satellite model.Satellites) error {
	_, err := r.collection.InsertOne(context.Background(), satellite)
	return err
}

// FindByName busca un documento por nombre.
func (r *SatelliteRepositoryMongo) FindByName(name string) (model.Satellites, error) {
	var satellite model.Satellites
	filter := bson.M{"name": name}
	err := r.collection.FindOne(context.Background(), filter).Decode(&satellite)
	if err != nil {
		return model.Satellites{}, err
	}
	return satellite, nil
}

// FindAll obtiene todos los documentos de la colección.
func (r *SatelliteRepositoryMongo) FindAll() ([]model.Satellites, error) {
	var satellites []model.Satellites

	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var satellite model.Satellites
		if err := cursor.Decode(&satellite); err != nil {
			return nil, err
		}
		satellites = append(satellites, satellite)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return satellites, nil
}

// FindByNames busca varios documentos por una lista de nombres.
func (r *SatelliteRepositoryMongo) FindByNames(names []string) ([]model.Satellites, error) {
	var satellites []model.Satellites
	filter := bson.M{"name": bson.M{"$in": names}}

	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var satellite model.Satellites
		if err := cursor.Decode(&satellite); err != nil {
			return nil, err
		}
		satellites = append(satellites, satellite)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return satellites, nil
}

// Delete elimina un documento por nombre.
func (r *SatelliteRepositoryMongo) Delete(name string) error {
	filter := bson.M{"name": name}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}

// DeleteAll elimina todos los documentos de la colección.
func (r *SatelliteRepositoryMongo) DeleteAll() error {
	_, err := r.collection.DeleteMany(context.Background(), bson.M{})
	return err
}

package awsSecret

import (
	"context"
	"encoding/json"
	"fuego-quasar-app/internal/core/domain/model"
	"fuego-quasar-app/internal/core/domain/port"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type AWSSecretManagerService struct{}

func NewAWSSecretManagerService() port.SecretManagerService {
	return &AWSSecretManagerService{}
}
func (t *AWSSecretManagerService) GetSecret(secretName string) (model.Setting, error) {
	region := "us-east-1"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatal(err.Error())
	}

	secretJson := *result.SecretString

	var setting model.Setting

	// Deserializar el JSON en la estructura
	err = json.Unmarshal([]byte(secretJson), &setting)

	if err != nil {
		log.Fatalf("Error retrieving secret: %v", err)
	}

	return setting, nil

}

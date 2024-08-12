require (
	github.com/aws/aws-lambda-go v1.36.1
	github.com/aws/aws-sdk-go v1.55.5
	github.com/aws/aws-sdk-go-v2 v1.30.3 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.27.27
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.32.4
	github.com/google/wire v0.6.0
	github.com/stretchr/testify v1.9.0
	go.mongodb.org/mongo-driver v1.16.1
)

replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8

module fuego-quasar-app

go 1.16

# Documentación del Proyecto AWS Lambda con SAM

## 1. Introducción

- **Nombre del Proyecto**: `fuego-quasar-app`
- **Descripción**: Este proyecto es una implementación en Go que utiliza la arquitectura hexagonal para calcular la ubicación de una nave a partir de distancias medidas por tres satélites y reconstruir el mensaje que la nave emite. El proyecto se despliega como una función Lambda utilizando AWS SAM.

## 2. Estructura del Proyecto

```plaintext
fuego-quasar-app/
├── internal/
│   ├── core/
│   │   ├── application/
│   │   │   └── service/
│   │   │       ├── decodeMessageService.go
│   │   │       ├── fuegoQuasarService.go
│   │   │       └── triangulationService.go
│   │   └── domain/
│   │       ├── model/
│   │       │   ├── point.go
│   │       │   ├── response.go
│   │       │   ├── satellite.go
│   │       │   └── setting.go
│   │       └── port/
│   │           ├── decodeMessageService.go
│   │           ├── fuegoQuasarService.go
│   │           ├── satelliteRepository.go
│   │           ├── secretManagerService.go
│   │           └── triangulationService.go
│   ├── infrastructure/
│   │   ├── awsSecret/
│   │   │   └── AWSSecretManagerService.go
│   │   ├── di/
│   │   │   ├── wire.go
│   │   │   └── wire_gen.go
│   │   ├── mongodb/
│   │   │   └── mongoClient.go
│   │   └── repository/
│   │       └── satelliteRepositoryMongo.go
│   └── interfaces/
│       └── handler/
│           └── lambdaHandler.go
├── test/
│   ├── decodeMessageService_test.go
│   └── triangulationService_test.go
├── go.mod
├── go.sum
├── main.go
└── main_test.go



```

La estructura del proyecto `fuego-quasar-app` está organizada de acuerdo con los principios de la arquitectura hexagonal. A continuación se describe cada capa y su propósito dentro del proyecto:

### 2.1 `internal/`

La carpeta `internal` contiene el código fuente del proyecto que está reservado para su uso interno y no debe ser importado desde otros proyectos o paquetes externos.

#### 2.1.1 `core/`

La capa `core` es el núcleo de la aplicación, donde se encuentra la lógica de negocio y las definiciones del dominio.

- **`application/`**: Contiene los servicios de la aplicación, que implementan la lógica de negocio y las operaciones del dominio.

  - **`service/`**: Implementa los servicios específicos de la aplicación.
    - `decodeMessageService.go`: Servicio responsable de decodificar mensajes, basado en el modelo de dominio.
    - `fuegoQuasarService.go`: Servicio específico del dominio, manejando la lógica central del proyecto.
    - `triangulationService.go`: Servicio para manejar la lógica de triangulación, utilizada para posicionamiento.

- **`domain/`**: Define los elementos del dominio de la aplicación, incluyendo modelos y puertos.

  - **`model/`**: Contiene las estructuras de datos utilizadas en el dominio.
    - `point.go`: Define estructuras y lógica para representar puntos en un espacio 2D o 3D.
    - `response.go`: Define las estructuras para las respuestas de la aplicación.
    - `satellite.go`: Define el modelo para los satélites.
    - `setting.go`: Contiene configuraciones y ajustes específicos del dominio.

  - **`port/`**: Define las interfaces que representan los puertos de entrada y salida de la aplicación. Estos puertos permiten la interacción con la lógica de negocio desde el exterior.
    - `decodeMessageService.go`: Define la interfaz para el servicio de decodificación de mensajes.
    - `fuegoQuasarService.go`: Define la interfaz para el servicio específico del dominio.
    - `satelliteRepository.go`: Define la interfaz para el repositorio de satélites.
    - `secretManagerService.go`: Define la interfaz para la gestión de secretos.
    - `triangulationService.go`: Define la interfaz para el servicio de triangulación.

#### 2.1.2 `infrastructure/`

La capa `infrastructure` contiene las implementaciones que interactúan con sistemas externos, como bases de datos y servicios externos.

- **`awsSecret/`**: Implementa la gestión de secretos utilizando AWS Secrets Manager.
  - `AWSSecretManagerService.go`: Implementa la interfaz para acceder y gestionar los secretos almacenados en AWS Secrets Manager.

- **`di/`**: Contiene la configuración para la inyección de dependencias.
  - `wire.go`: Define las dependencias e implementaciones necesarias utilizando el framework Wire.
  - `wire_gen.go`: Archivo generado automáticamente por Wire que contiene el código para la inyección de dependencias.


- **`mongodb/`**: Contiene la configuración y las implementaciones relacionadas con la base de datos **MongoDB**.
  - `mongoClient.go`: Configura y proporciona el cliente para interactuar con **MongoDB**.

- **`repository/`**: Implementaciones de los repositorios que interactúan con los sistemas de almacenamiento de datos.
  - `satelliteRepositoryMongo.go`: Implementa la interfaz del repositorio de satélites utilizando **MongoDB** como sistema de almacenamiento.

#### 2.1.3 `interfaces/`

La capa `interfaces` define los adaptadores que transforman las solicitudes y respuestas entre el mundo exterior y la lógica de negocio de la aplicación.

- **`handler/`**: Maneja las solicitudes y respuestas de la interfaz de la aplicación.
  - `lambdaHandler.go`: Adaptador para manejar las solicitudes provenientes de AWS Lambda, transformándolas en un formato que puede ser procesado por los servicios de la aplicación.

### 2.2 `test/`

Contiene las pruebas unitarias para asegurar que la lógica de la aplicación funcione correctamente.

- `decodeMessageService_test.go`: Pruebas unitarias para el servicio de decodificación de mensajes.
- `triangulationService_test.go`: Pruebas unitarias para el servicio de triangulación.

### 2.3 Inyección de dependencias co Wire

Wire es una herramienta para la inyección de dependencias en Go, creada por Google. Facilita la configuración automática de dependencias y la gestión de la inyección de dependencias en proyectos complejos. Aquí se describe cómo se integra Wire en el proyecto:

- **`wire.go`**: Este archivo define los proveedores y las dependencias necesarias para el proyecto. Utiliza las anotaciones de Wire para especificar cómo se deben construir y conectar las dependencias. Este archivo debe contener funciones que definan la creación de los objetos y su configuración.

- **`wire_gen.go`**: Archivo generado automáticamente por Wire. Contiene el código que Wire genera en base a las configuraciones de `wire.go`. No debes modificar este archivo manualmente; Wire lo actualiza cuando ejecutas el comando de generación.

#### 2.3.1  Ejemplo de Uso de Wire

##### 2.3.1.1 **Definir Proveedores en `wire.go`**:

   ```go
   ///go:build wireinject
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

   ```
##### 2.3.1.2 Generar el Código de Inyección de Dependencias:

Ejecuta el siguiente comando para generar el archivo wire_gen.go:

```go
wire ./internal/infrastructure/di
```
Esta integración con Wire ayuda a simplificar la gestión de dependencias y mejora la mantenibilidad del código en proyectos grandes.

## 3. Configuración de AWS SAM

### 3.1. Archivo `template.yaml`

El archivo `template.yaml` es el archivo principal de configuración para AWS SAM. Aquí está un ejemplo de cómo se vería este archivo:

```yaml
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31
Description: >-
  Fuego Quasar App - Lambda Function
Globals:
  Function:
    Timeout: 10
    MemorySize: 128

Resources:
  FuegoQuasarFunction:
    Type: AWS::Serverless::Function
    Properties:
      Handler: cmd/lambda/main
      Runtime: go1.x
      CodeUri: ./
      Environment:
        Variables:
          CONNECTION_STRING_SETTING: "prod/conectionstringfuegoquasardb"
      Events:
        FuegoQuasarApi:
          Type: Api
          Properties:
            Path: /api/v1/satellites
            Method: post
```
### 3.2. Variables de Entorno
Estas son las variables de entorno que usa la app para su correcto funcionamiento.
- **`CONNECTION_SECRET_NAME: prod/connectionstringfuegoquasardb`** cadena con el nombre del secreto que tiene la cadena de conexion a **MongoDB**
- **`KENOBI_X: -500`** posicion X del satelite kenobi
- **`KENOBI_Y: -200`** posicion Y del satelite kenobi
- **`SKYWALKER_X: 100`** posicion X del satelite skywalker
- **`SKYWALKER_Y: -100`** posicion Y del satelite skywalker
- **`SATO_X: 500`** posicion X del satelite sato
- **`SATO_Y: 100`** posicion Y del satelite sato




### 3.3. Despliegue con AWS SAM

#### 3.3.1. Instalación de AWS SAM CLI

Asegúrate de tener instalado AWS SAM CLI. Si no lo tienes, puedes instalarlo siguiendo las instrucciones en la [documentación oficial de AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html).

#### 3.3.2. Compilación del Proyecto

Antes de desplegar la función, debes compilar el proyecto. Ejecuta el siguiente comando en el directorio raíz del proyecto:

```sh
sam build

```
Este comando te guiará a través del proceso de despliegue, donde deberás proporcionar un nombre para el stack de CloudFormation y otros parámetros de configuración.


## 4. Ejemplos de Uso

### 4.1. Invocar la Función Lambda

Puedes invocar la función Lambda utilizando la AWS CLI o una herramienta como Postman. Aquí hay un ejemplo utilizando la AWS CLI:

```sh
aws lambda invoke \
    --function-name FuegoQuasarFunction \
    --payload file://input.json \
    output.json
```
## 5. Pruebas Unitarias

### 5.1. Ejecutar todas las pruebas

Para ejecutar todas las pruebas unitarias del proyecto, usa el siguiente comando:

```sh
go test ./...
```

### 5.2. Ejecutar Pruebas con Cobertura

Para ejecutar las pruebas unitarias y generar un informe de cobertura, utiliza el siguiente comando:

```sh
go test -coverprofile=coverage.out ./...
```

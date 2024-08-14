# Documentación del Proyecto AWS Lambda con SAM

## 1. Introducción

- **Nombre del Proyecto**: `fuego-quasar-app`
- **Descripción**: Este proyecto es una implementación en Go que utiliza la arquitectura hexagonal para calcular la ubicación de una nave a partir de distancias medidas por tres satélites y reconstruir el mensaje que la nave emite. El proyecto se despliega como una función Lambda utilizando AWS SAM.

### 1.1 Solución al problema de encontrar al punto
En un problema de trilateración con tres puntos de referencia en un plano 2D y un cuarto punto desconocido cuya distancia a estos tres puntos es conocida, queremos encontrar las coordenadas del punto desconocido.

#### 1.1.2 Sistema de Ecuaciones

Dado tres puntos de referencia \((x_1, y_1)\), \((x_2, y_2)\), y \((x_3, y_3)\), y las distancias desde el punto desconocido \((x, y)\) a estos puntos \(d_1\), \(d_2\), y \(d_3\), las ecuaciones basadas en la distancia euclidiana son:

1. 
   ![Ecuación 1](https://latex.codecogs.com/gif.latex?(x%20-%20x_1)^2%20%2B%20(y%20-%20y_1)^2%20%3D%20d_1^2)

2.
   ![Ecuación 2](https://latex.codecogs.com/gif.latex?(x%20-%20x_2)^2%20%2B%20(y%20-%20y_2)^2%20%3D%20d_2^2)

3. 
   ![Ecuación 3](https://latex.codecogs.com/gif.latex?(x%20-%20x_3)^2%20%2B%20(y%20-%20y_3)^2%20%3D%20d_3^2)

la solucion de este sistema de ecuaciones nos da como resultado el X, Y

#### 1.1.3  Verificación Algebraica

##### 1.1.3.1  Cálculo del Determinante

Para verificar la existencia y unicidad de la solución, se debe evaluar el determinante del sistema lineal obtenido al restar pares de ecuaciones cuadráticas. El determinante ayuda a determinar si el sistema es resoluble.

- **Determinante**:


![Determinador](https://latex.codecogs.com/gif.latex?denominator%20%3D%202(x_2%20-%20x_1)%20*%202(y_3%20-%20y_1)%20-%202(y_2%20-%20y_1)%20*%202(x_3%20-%20x_1))

Si el denominador es cero, las ecuaciones pueden ser linealmente dependientes, lo que puede indicar que el sistema no tiene una solución única. En este caso, verifica si las ecuaciones son inconsistentes o si el sistema tiene soluciones infinitas.

##### 1.1.3.1 Solución del Sistema Lineal

Resuelve el sistema lineal para las coordenadas \(x\) e \(y\) usando las siguientes fórmulas:

- **Solución para \(x\)**:


![Solución X](https://latex.codecogs.com/gif.latex?x%20%3D%20%5Cfrac%7B(d_1^2%20-%20d_2^2%20%2B%20x_2^2%20-%20x_1^2%20%2B%20y_2^2%20-%20y_1^2)%20*%202(y_3%20-%20y_1)%20-%20(d_1^2%20-%20d_3^2%20%2B%20x_3^2%20-%20x_1^2%20%2B%20y_3^2%20-%20y_1^2)%20*%202(y_2%20-%20y_1)%7D%7Bdenominator%7D)

- **Solución para \(y\)**:


![Solución Y](https://latex.codecogs.com/gif.latex?y%20%3D%20%5Cfrac%7B(d_1^2%20-%20d_2^2%20%2B%20x_2^2%20-%20x_1^2%20%2B%20y_2^2%20-%20y_1^2)%20*%202(x_1%20-%20x_2)%20-%20(d_1^2%20-%20d_3^2%20%2B%20x_1^2%20-%20x_3^2%20%2B%20y_1^2%20-%20y_3^2)%20*%202(x_1%20-%20x_3)%7D%7Bdenominator%7D)

##### 1.1.3. Conclusión
**Solución Única**: Si el determinante no es cero y las distancias cumplen las condiciones triangulares, hay una solución única para  (𝑥,𝑦)

**No hay Solución**: Si el determinante es cero y las distancias no cumplen las condiciones triangulares, o si las ecuaciones son inconsistentes, no hay solución válida.

**Soluciones Múltiples**: Si el determinante es cero pero las ecuaciones son consistentes, puede haber soluciones infinitas o ninguna solución dependiendo de las condiciones adicionales.

Estas verificaciones aseguran que el sistema de ecuaciones tiene una solución válida y ayuda a identificar posibles problemas en los datos o en la implementación del algoritmo.

### 1.2 Solución al problema de decodificar el mensaje

El paquete `service` proporciona una implementación para decodificar mensajes a partir de un conjunto de datos de entrada. Esta implementación se basa en la idea de que cada entrada en el mensaje puede tener palabras en una posición específica, y el objetivo es construir un mensaje a partir de la palabra más frecuente en cada posición.

### DecodeMessageService

Esta funcion de servicio está diseñado para decodificar mensajes a partir de una matriz de cadenas.

#### Métodos

- **GetMessage(message [][]string) (string, error)**: Decodifica el mensaje dado. Combina las palabras más frecuentes en cada posición de las sublistas del mensaje para construir la cadena final. Retorna el mensaje decodificado o un error si hay problemas con la longitud del mensaje o si el resultado está vacío.

## Funciones Auxiliares

### getMessageLength

Calcula la longitud máxima del mensaje basada en el tamaño de las sublistas.

### getWordByPosition

Obtiene la palabra más frecuente en una posición específica de las sublistas.

### deleteOffset

Elimina los elementos anteriores a una longitud específica de cada sublista en el mensaje.

### getMessageLengthFirtsWord

Encuentra la palabra más frecuente en la primera posición y devuelve su índice y la longitud de la sublista correspondiente.

### removeEmptyStrings

Elimina las cadenas vacías de una lista de strings.
En resumen, este servicio toma un conjunto de datos en forma de matriz de cadenas, encuentra la palabra más frecuente en cada posición, y construye el mensaje decodificado final. Además, maneja errores relacionados con la longitud del mensaje y el contenido resultante.

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
![Arquitectura del Core del Proyecto](img/core.png?raw=true "Diagrama de la arquitectura core del proyecto")
*Diagrama de la arquitectura core del proyecto*

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
![Arquitectura del Infraestructura del Proyecto](img/infraestructura.png?raw=true "Diagrama de la arquitectura core del proyecto")
*Diagrama de la arquitectura infraestructura del proyecto*

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
Description: >
  fuego-quasar-meli-app

  Plantilla SAM  para la función fuego-quasar-meli-app

# More info about Globals: https://github.com/awslabs/serverless-application-model/blob/master/docs/globals.rst
Globals:
  Function:
    Timeout: 5
    MemorySize: 128

    Tracing: Active
    # You can add LoggingConfig parameters such as the Logformat, Log Group, and SystemLogLevel or ApplicationLogLevel. Learn more here https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-resource-function.html#sam-function-loggingconfig.
    LoggingConfig:
      LogFormat: JSON
  Api:
    TracingEnabled: true
Resources:
  FuegoQuasarFunction:
    Type: AWS::Serverless::Function # More info about Function Resource: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
    Metadata:
      BuildMethod: go1.x
    Properties:
      CodeUri: fuego-quasar-app/
      Handler: bootstrap
      Runtime: provided.al2023
      Architectures:
      - x86_64
      Events:
        Topsecret:
          Type: Api # More info about API Event Source: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
          Properties:
            Path: /topsecret
            Method: POST
        PosttopsecretSplit:
          Type: Api # More info about API Event Source: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
          Properties:
            Path: /topsecret_split
            Method: POST
        GettopsecretSplit:
          Type: Api # More info about API Event Source: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
          Properties:
            Path: /topsecret_split
            Method: GET
      Environment: # More info about Env Vars: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#environment-object
        Variables:
          CONNECTION_SECRET_NAME: prod/connectionstringfuegoquasardb
          KENOBI_X: -500
          KENOBI_Y: -200
          SKYWALKER_X: 100
          SKYWALKER_Y: -100
          SATO_X: 500
          SATO_Y: 100



  ApplicationResourceGroup:
    Type: AWS::ResourceGroups::Group
    Properties:
      Name:
        Fn::Sub: ApplicationInsights-SAM-${AWS::StackName}
      ResourceQuery:
        Type: CLOUDFORMATION_STACK_1_0
  ApplicationInsightsMonitoring:
    Type: AWS::ApplicationInsights::Application
    Properties:
      ResourceGroupName:
        Ref: ApplicationResourceGroup
      AutoConfigurationEnabled: 'true'
Outputs:
  # ServerlessRestApi is an implicit API created out of Events key under Serverless::Function
  # Find out more about other implicit resources you can reference within SAM
  # https://github.com/awslabs/serverless-application-model/blob/master/docs/internals/generated_resources.rst#api
  FuegoQuasardAPI:
    Description: API Gateway endpoint URL for Prod environment for First Function
    Value: !Sub "https://${ServerlessRestApi}.execute-api.${AWS::Region}.amazonaws.com/Prod/topsecret/"
  FuegoQuasardAPI2:
    Description: API Gateway endpoint URL for Prod environment for First Function
    Value: !Sub "https://${ServerlessRestApi}.execute-api.${AWS::Region}.amazonaws.com/Prod/topsecret_split/"
  FuegoQuasarFunction:
    Description: First Lambda Function ARN
    Value: !GetAtt FuegoQuasarFunction.Arn
  FuegoQuasarFunctionIamRole:
    Description: Implicit IAM Role created for Fuego Quasarfunction
    Value: !GetAtt FuegoQuasarFunctionRole.Arn

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
## 5. Pruebas

### 5.1. Pruebas Unitarias

Para ejecutar todas las pruebas unitarias del proyecto, usa el siguiente comando:

```sh
go test ./...
```
### 5.2 Pruebas de regresión
Para asegurar la calidad y el estado del servicio se crean pruebas de consumo del api, teniendo en cuenta los diferentes casos de prueba expuestos en el reto propuesto, estas pruebas serán ejecutadas automaticamente en los procesos de CI/CD.
Se pude consultar esta **[Coleccion Postam](https://lunar-sunset-766256.postman.co/workspace/Meli_Fire_Quasar~bd17065b-4543-4236-923b-8781260d6a56/collection/2242228-bbadee89-3d1d-434b-bb27-4220d2738fda?action=share&creator=2242228&active-environment=2242228-c25adb19-d282-4e15-ba47-4a3fec9549c3)**

## 6 Proceso Global de CI/CD con AWS SAM y GitHub Actions

### 6.1  Descripción Global del Proceso de CI/CD

#### 6.1.1. Desencadenamiento del Proceso

El proceso de CI/CD se inicia automáticamente cuando hay un **push** a la rama `master` del repositorio en GitHub. Este proceso se encarga de construir y desplegar la aplicación en la infraestructura **AWS**, posterior a este proceso el sistema obtiene y ejecuta las pruebas necesarias para verificar el correcto funcionamiento de la aplicacion generando estos reportes de la ejecución de las mismas.

![Pruebas de regresión](img/pruebasRegresion.png?raw=true "Pruebas de regresión")
Pruebas de regresión implementadas con postman



#### 6.1.2. Compilación y Despliegue (CI)

- **GitHub Actions** se utiliza como la plataforma de automatización que orquesta todo el flujo de trabajo de CI/CD.
- El código se **clona** del repositorio usando la acción `actions/checkout`.
- Las **credenciales de AWS** se configuran utilizando `aws-actions/configure-aws-credentials` para permitir que GitHub Actions interactúe con los servicios de AWS.
- La aplicación se **compila** utilizando el comando `sam build` de **AWS SAM** (Serverless Application Model).
- Después de la compilación, la aplicación se **despliega** en **Amazon ECS** utilizando el comando `sam deploy`.

#### 6.1.3. Pruebas (CD)

- Después del despliegue, se instala **Node.js** en el ambiente de GitHub Actions para ejecutar herramientas basadas en Node.
- Se instala **Newman**, el cliente de línea de comandos para ejecutar pruebas de API de Postman.
- Las **pruebas de API** se ejecutan usando Newman para verificar que la aplicación funciona como se espera.
- Los resultados de las pruebas se **suben** al repositorio de GitHub como artefactos para su revisión y análisis posterior.

cuando se ejecutan las pruebas usando la herramienta **newman** nos genera en los logs de la ejecuión del pipeline los resultados de las pruebas ejecutadas como se muestra a continuación.
![Ejecucion de pruebas de regresión](img/ejecucionPruebaCICD.png?raw=true "Ejecuion de las pruebas de regresión")
Ejecuión de las pruebas de regresión

ademas de esto **newman** genera un reporte detallado con el resultado de cada ejecuion y cada caso de pruba. aca podemos ver un ejemplo de un reporte generado por el proceso de **CI/CD** 

**[Reporte de pruebas](https://drive.usercontent.google.com/u/0/uc?id=1xh5yr4GpyYdiO1PcQXAFqogK7phuOHZ1&export=download)**

### 6.2 Herramientas Utilizadas

#### 6.2.1  GitHub Actions

- **Descripción**: Plataforma de automatización CI/CD que permite crear flujos de trabajo personalizados para compilar, probar y desplegar código directamente desde GitHub.
- **Función**: Orquesta todo el proceso de CI/CD, desde la compilación hasta el despliegue y pruebas.

#### 6.2.2 AWS SAM (Serverless Application Model)

- **Descripción**: Framework de código abierto para crear aplicaciones serverless en AWS. Simplifica el proceso de definición y despliegue de recursos en AWS.
- **Función**: Se utiliza para compilar y desplegar la aplicación Go en un entorno de AWS, permitiendo un despliegue fácil y automatizado.

#### 6.2.3 Newman (CLI de Postman)

- **Descripción**: Cliente de línea de comandos para ejecutar colecciones de pruebas de API de Postman desde cualquier lugar.
- **Función**: Ejecuta las pruebas de API de Postman para verificar la funcionalidad de la aplicación desplegada en AWS.

#### 6.2.4 Node.js

- **Descripción**: Entorno de ejecución de JavaScript que permite ejecutar código JavaScript en el lado del servidor.
- **Función**: Necesario para instalar y ejecutar Newman, que está escrito en JavaScript.

#### 6.2.5 Secrets de GitHub

- **Descripción**: Funcionalidad de GitHub que permite almacenar de forma segura las claves y credenciales necesarias para la autenticación y acceso a recursos externos.
- **Función**: Almacena credenciales de AWS y claves API de Postman, utilizadas durante el proceso de CI/CD.

## 7  Documentación de Seguridad del API con AWS Signature Version 4

### 7.1 Introducción

La seguridad de una API es esencial para proteger los datos y servicios que ofrece. AWS proporciona un mecanismo de autenticación robusto conocido como **AWS Signature Version 4 (SigV4)**, que se utiliza para autenticar y autorizar solicitudes a las APIs en AWS. Este método garantiza que las solicitudes provengan de entidades autorizadas y que los datos en tránsito no hayan sido manipulados.

### 7.2  Cómo Funciona la Autenticación con SigV4

El proceso de autenticación con SigV4 sigue los siguientes pasos:

1. **Recopilar Información de la Solicitud**: Incluye el método HTTP, la URL, los encabezados HTTP, y el cuerpo de la solicitud.
2. **Crear una Solicitud Canonical**: Formatear la información recopilada en una solicitud estructurada y normalizada.
3. **Crear la String to Sign**: Combinar la solicitud canonical con otros datos de la solicitud para crear una cadena de texto única.
4. **Derivar la Clave de Firma**: Utilizar la Secret Key del usuario de IAM para generar una clave de firma única.
5. **Crear la Firma**: Utilizar la clave de firma para crear una firma criptográfica de la string to sign.
6. **Agregar la Firma a la Solicitud**: Incluir la firma en el encabezado `Authorization` de la solicitud HTTP.

### 7.3 Formato del Encabezado `Authorization`

```plaintext
Authorization: AWS4-HMAC-SHA256 Credential=<Access Key ID>/<Date>/<Region>/<Service>/aws4_request, SignedHeaders=<Signed Headers>, Signature=<Signature>
```

### 7.4 Implementación en Postman

Para implementar este proceso en Postman:

1. **Abrir Postman**: Iniciar la aplicación Postman.
2. **Crear una Nueva Solicitud**: Hacer clic en "New" y seleccionar "Request".
3. **Configurar la Solicitud**: Ingresar la URL de la API y seleccionar el método HTTP apropiado.
4. **Agregar Encabezados**: En la sección "Headers", agregar los encabezados necesarios (`host`, `x-amz-date`, etc.).
5. **Firmar la Solicitud**: Antes de enviar la solicitud, usar una herramienta o script para generar el encabezado `Authorization` siguiendo los pasos descritos anteriormente.
6. **Enviar la Solicitud**: Hacer clic en "Send" para enviar la solicitud a la API.

### 7.5 Ventajas de Usar AWS Signature Version 4

- **Seguridad Mejorada**: Asegura que solo los usuarios autorizados puedan acceder a los recursos.
- **Integridad de Datos**: Garantiza que los datos no sean manipulados durante el tránsito.
- **Flexibilidad**: Funciona con una variedad de métodos HTTP y tipos de solicitudes.
- **Compatibilidad**: Es compatible con todos los servicios de AWS que soportan autenticación mediante SigV4.

### 7.6 AWS Identity and Access Management (IAM)

#### ¿Qué es IAM?

AWS Identity and Access Management (IAM) es un servicio que permite administrar el acceso a los recursos de AWS de manera segura. Con IAM, puedes:

- **Crear Usuarios y Roles**: Definir identidades que puedan autenticarse y realizar acciones en AWS.
- **Administrar Permisos**: Asignar políticas que especifiquen qué acciones puede realizar una identidad y en qué recursos.
- **Configurar Autenticación Multi-Factor (MFA)**: Añadir una capa extra de seguridad mediante la configuración de autenticación de dos factores.





package handler

import (
	"encoding/json"
	"fuego-quasar-app/internal/core/domain/model"
	"fuego-quasar-app/internal/core/domain/port"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaHandler struct {
	triangulationService port.TriangulationService
	decodeMessageService port.DecodeMessageService
	secretManagerService port.SecretManagerService
	satelliteRepository  port.SatelliteRepository
	fuegoQuasarService   port.FuegoQuasarService
	logService           port.LogService
}

func NewLambdaHandler(logService port.LogService, triangulationService port.TriangulationService, decodeMessageService port.DecodeMessageService, secretManagerService port.SecretManagerService, satelliteRepository port.SatelliteRepository, fuegoQuasarService port.FuegoQuasarService) LambdaHandler {
	return LambdaHandler{logService: logService, triangulationService: triangulationService, decodeMessageService: decodeMessageService, secretManagerService: secretManagerService, satelliteRepository: satelliteRepository, fuegoQuasarService: fuegoQuasarService}
}

func (h LambdaHandler) HandlePostRequestTopsecret_split(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	body := request.Body
	var satellite model.Satellites

	// Obtener el cuerpo de la solicitud

	// Deserializar el cuerpo JSON a la estructura RequestBody
	err := json.Unmarshal([]byte(body), &satellite)
	if err != nil {
		h.logService.Info("HandleRequest:", "error", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error al procesar el cuerpo de la solicitud. Daat",
		}, nil
	}
	err = h.fuegoQuasarService.ProcessSplitMessage(satellite)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 400,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       "OK",
		StatusCode: 200,
	}, nil
}

func (h LambdaHandler) HandleGetRequestTopsecret_split(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	response, err := h.fuegoQuasarService.ProcessSaveMessages()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 400,
		}, nil

	}
	// Serializar la estructura a JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 400,
		}, nil

	}
	// Convertir []byte a string para imprimirlo
	jsonString := string(jsonData)

	return events.APIGatewayProxyResponse{
		Body:       jsonString,
		StatusCode: 200,
	}, nil
}

func (h LambdaHandler) HandlePostRequestTopsecret(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	body := request.Body
	var satellites []model.Satellites

	// Obtener el cuerpo de la solicitud
	satellites = nil
	// Deserializar el cuerpo JSON a la estructura RequestBody
	err := json.Unmarshal([]byte(body), &satellites)
	if err != nil {
		h.logService.Info("HandleRequest:", "error", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error al procesar el cuerpo de la solicitud.",
		}, nil
	}
	response, err := h.fuegoQuasarService.ProcessMessages(satellites)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 400,
		}, nil
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 400,
		}, nil

	}

	// Convertir []byte a string para imprimirlo
	jsonString := string(jsonData)

	return events.APIGatewayProxyResponse{
		Body:       jsonString,
		StatusCode: 200,
	}, nil
}

func (h LambdaHandler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	h.logService.Info("HandleRequest:", "request", request)

	switch request.Path {
	case "/topsecret":
		if request.HTTPMethod == "POST" {
			return h.HandlePostRequestTopsecret(request)
		}
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       "Método no permitido",
		}, nil

	case "/topsecret_split":
		switch request.HTTPMethod {
		case "GET":
			return h.HandleGetRequestTopsecret_split(request)
		case "POST":
			return h.HandlePostRequestTopsecret_split(request)
		default:
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusMethodNotAllowed,
				Body:       "Método no permitido",
			}, nil
		}
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "Endpoint no encontrado",
		}, nil
	}

}

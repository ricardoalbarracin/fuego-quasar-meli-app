package handler

import (
	"fmt"
	"fuego-quasar-app/internal/core/domain/model"
	"fuego-quasar-app/internal/core/domain/port"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaHandler struct {
	triangulationService port.TriangulationService
	decodeMessageService port.DecodeMessageService
	secretManagerService port.SecretManagerService
	satelliteRepository  port.SatelliteRepository
	fuegoQuasarService   port.FuegoQuasarService
}

func NewLambdaHandler(triangulationService port.TriangulationService, decodeMessageService port.DecodeMessageService, secretManagerService port.SecretManagerService, satelliteRepository port.SatelliteRepository, fuegoQuasarService port.FuegoQuasarService) LambdaHandler {
	return LambdaHandler{triangulationService: triangulationService, decodeMessageService: decodeMessageService, secretManagerService: secretManagerService, satelliteRepository: satelliteRepository, fuegoQuasarService: fuegoQuasarService}
}

func (h *LambdaHandler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//aaa, _ := h.satelliteRepository.FindByName("DummySatellite")
	//fmt.Println("Secret:", aaa.Name)
	h.satelliteRepository.DeleteAll()
	h.satelliteRepository.Create(&model.Satellites{
		Name:     "kenobi",
		Distance: 100.0,
		Message:  []string{"este", "", "", "mensaje", ""},
	})

	h.satelliteRepository.Create(&model.Satellites{
		Name:     "skywalker",
		Distance: 115.5,
		Message:  []string{"", "es", "", "", "secreto"},
	})

	h.satelliteRepository.Create(&model.Satellites{
		Name:     "sato",
		Distance: 142.7,
		Message:  []string{"este", "", "un", "", ""},
	})
	msg, _ := h.fuegoQuasarService.ProcessSaveMessages()
	fmt.Printf("XX: %f YY: %f\n", msg.Position.X, msg.Position.Y)

	fmt.Printf(" pro: %s\n", msg.Message)
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP
	//secretName := "prod/conectionstringfuegoquasardb"
	//secret, _ := h.secretManagerService.GetSecret(secretName)

	satellites, _ := h.satelliteRepository.FindAll()
	for _, satellite := range satellites {
		fmt.Printf("Satellite Name: %s\n", satellite.Name)
		fmt.Printf("Distance: %f\n", satellite.Distance)
		fmt.Printf("Message: %v\n", satellite.Message)
		fmt.Println("------------------------------")
	}

	p1 := model.Point{X: -500, Y: -200}
	p2 := model.Point{X: 100, Y: -100}
	p3 := model.Point{X: 500, Y: 100}

	// Distancias desde el dispositivo hasta los puntos de referencia
	d1 := 100.0
	d2 := 115.5
	d3 := 142.7

	listOfLists := [][]string{
		{"", "este", "es", "un", "mensaje"},
		{"este", "", "un", "mensaje"},
		{"", "este", "es", "", ""},
	}
	// Calculo de la posicion
	punto, _ := h.triangulationService.GetLocation(p1, p2, p3, d1, d2, d3)
	fmt.Printf("X: %f Y: %f", punto.X, punto.Y)
	mensaje, _ := h.decodeMessageService.GetMessage(listOfLists)
	fmt.Printf("mensaje: %s", mensaje)
	//myEnvVar := os.Getenv("PARAM1")
	if sourceIP == "" {
		greeting = fmt.Sprintf("HellXXasaYYYYXXXXSo, X: %f Y: %f\n", punto.X, punto.Y)
	} else {
		greeting = fmt.Sprintf("HellXXXXXXSo, %f!\n", 1)
	}

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: 200,
	}, nil
}

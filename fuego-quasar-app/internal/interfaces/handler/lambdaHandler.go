package handler

import (
	"fmt"
	"fuego-quasar-app/internal/core/domain/model"
	"fuego-quasar-app/internal/core/domain/port"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaHandler struct {
	triangulationService port.TriangulationService
}

func NewLambdaHandler(userService port.TriangulationService) LambdaHandler {
	return LambdaHandler{triangulationService: userService}
}

func (h *LambdaHandler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	p1 := model.Point{X: 0, Y: 0}
	p2 := model.Point{X: 4, Y: 0}
	p3 := model.Point{X: 2, Y: 4}

	// Distancias desde el dispositivo hasta los puntos de referencia
	d1 := 2.0
	d2 := 2.828
	d3 := 2.828

	// Calculo de la posicion
	punto, errores := h.triangulationService.GetLocation(p1, p2, p3, d1, d2, d3)

	if sourceIP == "" {
		greeting = fmt.Sprintf("HellXXasaYYYYXXXXSo, %f! %f! \n", punto.X, punto.Y)
	} else {
		greeting = fmt.Sprintf("HellXXXXXXSo, %f!\n", punto.X)
	}

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: 200,
	}, errores
}

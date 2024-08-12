package service

import (
	"fuego-quasar-app/internal/core/domain/model"
	"fuego-quasar-app/internal/core/domain/port"
	"log"
	"os"
	"strconv"
)

type FuegoQuasarService struct {
	satelliteRepository  port.SatelliteRepository
	decodeMessageService port.DecodeMessageService
	triangulationService port.TriangulationService
}

func NewFuegoQuasarService(satelliteRepository port.SatelliteRepository, decodeMessageService port.DecodeMessageService, triangulationService port.TriangulationService) port.FuegoQuasarService {
	return &FuegoQuasarService{satelliteRepository: satelliteRepository, decodeMessageService: decodeMessageService, triangulationService: triangulationService}
}

func (t *FuegoQuasarService) ProcessSplitMessage(satellite model.Satellites) error {

	t.satelliteRepository.Delete(satellite.Name)
	err := t.satelliteRepository.Create(&satellite)

	return err
}

func (t *FuegoQuasarService) ProcessSaveMessages() (model.Response, error) {

	namesSatellites := []string{"kenobi", "skywalker", "sato"}
	satellites, _ := t.satelliteRepository.FindByNames(namesSatellites)

	if len(satellites) < 3 {
		log.Fatalf("there is not enough information")
	}
	var messages [][]string
	for _, satellite := range satellites {
		messages = append(messages, satellite.Message)

	}
	message, err := t.decodeMessageService.GetMessage(messages)
	if err != nil {
		log.Fatalf("Error decode Message: %v", err)
	}
	p1, p2, p3, err := t.GetSatellitesPoints()
	if err != nil {
		log.Fatalf("Error get satellites points: %v", err)
	}
	d1, err := t.FilterSatelliteByName(satellites, "kenobi")
	if err != nil {
		log.Fatalf("Error get satellites distance: %v", err)
	}
	d2, err := t.FilterSatelliteByName(satellites, "skywalker")
	if err != nil {
		log.Fatalf("Error get satellites distance: %v", err)
	}
	d3, err := t.FilterSatelliteByName(satellites, "sato")
	if err != nil {
		log.Fatalf("Error get satellites distance: %v", err)
	}
	position, err := t.triangulationService.GetLocation(p1, p2, p3, d1.Distance, d2.Distance, d3.Distance)

	if err != nil {
		log.Fatalf("Error get satellites GetLocation: %v", err)
	}
	return model.Response{Message: message, Position: position}, nil
}

func (t *FuegoQuasarService) FilterSatelliteByName(satellites []*model.Satellites, name string) (*model.Satellites, error) {
	for _, satellite := range satellites {
		if satellite.Name == name {
			return satellite, nil
		}
	}
	log.Fatalf("Error satellite not found")
	return nil, nil
}

func (t *FuegoQuasarService) GetSatellitesPoints() (model.Point, model.Point, model.Point, error) {
	kenobiX, err := strconv.ParseFloat(os.Getenv("KENOBI_X"), 64)
	if err != nil {
		log.Fatalf("Error conver to float: %v", err)
	}
	kenobiY, err := strconv.ParseFloat(os.Getenv("KENOBI_Y"), 64)
	if err != nil {
		log.Fatalf("Error conver to float: %v", err)
	}

	skywalkerX, err := strconv.ParseFloat(os.Getenv("SKYWALKER_X"), 64)
	if err != nil {
		log.Fatalf("Error conver to float: %v", err)
	}
	skywalkerY, err := strconv.ParseFloat(os.Getenv("SKYWALKER_Y"), 64)
	if err != nil {
		log.Fatalf("Error conver to float: %v", err)
	}

	satoX, err := strconv.ParseFloat(os.Getenv("SATO_X"), 64)
	if err != nil {
		log.Fatalf("Error conver to float: %v", err)
	}
	satoY, err := strconv.ParseFloat(os.Getenv("SATO_Y"), 64)
	if err != nil {
		log.Fatalf("Error conver to float: %v", err)
	}

	return model.Point{X: kenobiX, Y: kenobiY}, model.Point{X: skywalkerX, Y: skywalkerY}, model.Point{X: satoX, Y: satoY}, nil
}

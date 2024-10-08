package service

import (
	"errors"
	"fmt"
	"fuego-quasar-app/internal/core/domain/model"
	"fuego-quasar-app/internal/core/domain/port"
	"os"
	"strconv"
	"time"
)

type FuegoQuasarService struct {
	satelliteRepository  port.SatelliteRepository
	decodeMessageService port.DecodeMessageService
	triangulationService port.TriangulationService
	logService           port.LogService
}

func NewFuegoQuasarService(logService port.LogService, satelliteRepository port.SatelliteRepository, decodeMessageService port.DecodeMessageService, triangulationService port.TriangulationService) port.FuegoQuasarService {
	return FuegoQuasarService{logService: logService, satelliteRepository: satelliteRepository, decodeMessageService: decodeMessageService, triangulationService: triangulationService}
}

func (t FuegoQuasarService) ProcessSplitMessage(satellite model.Satellites) error {
	t.satelliteRepository.Delete(satellite.Name)
	time.Sleep(100 * time.Millisecond)
	err := t.satelliteRepository.Create(satellite)

	return err
}

func (t FuegoQuasarService) ProcessSaveMessages() (model.Response, error) {

	t.logService.Info("ProcessSaveMessages")
	namesSatellites := []string{"kenobi", "skywalker", "sato"}
	satellites, _ := t.satelliteRepository.FindByNames(namesSatellites)

	numMessage := len(satellites)
	if numMessage < 3 {
		t.logService.Error("no hay suficiente información", "satellites", satellites)
		errorMessage := fmt.Sprintf("no hay suficiente información %d", numMessage)
		return model.Response{}, errors.New(errorMessage)
	}
	var messages [][]string
	for _, satellite := range satellites {
		messages = append(messages, satellite.Message)

	}
	message, err := t.decodeMessageService.GetMessage(messages)
	if err != nil {
		t.logService.Error("error decodeMessageService", "Error", err)
		return model.Response{}, err
	}
	p1, p2, p3, err := t.GetSatellitesPoints()
	if err != nil {
		t.logService.Error("error GetSatellitesPoints", "Error", err)
		return model.Response{}, err
	}
	d1, err := t.FilterSatelliteByName(satellites, "kenobi")
	if err != nil {
		t.logService.Error("error FilterSatelliteByName kenobi", "Error", err)
		return model.Response{}, err
	}
	d2, err := t.FilterSatelliteByName(satellites, "skywalker")
	if err != nil {
		t.logService.Error("error FilterSatelliteByName skywalker", "Error", err)
		return model.Response{}, err
	}
	d3, err := t.FilterSatelliteByName(satellites, "sato")
	if err != nil {

		t.logService.Error("error FilterSatelliteByName sato", "Error", err)
		return model.Response{}, err
	}
	position, err := t.triangulationService.GetLocation(p1, p2, p3, d1.Distance, d2.Distance, d3.Distance)

	t.satelliteRepository.DeleteAll()
	if err != nil {
		return model.Response{}, err
	}
	return model.Response{Message: message, Position: position}, nil
}

func (t FuegoQuasarService) ProcessMessages(satellites []model.Satellites) (model.Response, error) {

	if len(satellites) < 3 {
		t.logService.Error("error o hay suficiente información", "Error", "satellites", satellites)
		return model.Response{}, errors.New("no hay suficiente información")
	}
	var messages [][]string
	for _, satellite := range satellites {
		messages = append(messages, satellite.Message)
	}

	message, err := t.decodeMessageService.GetMessage(messages)
	if err != nil {
		return model.Response{}, err
	}
	p1, p2, p3, err := t.GetSatellitesPoints()
	if err != nil {
		return model.Response{}, err
	}
	d1, err := t.FilterSatelliteByName(satellites, "kenobi")
	if err != nil {
		return model.Response{}, err
	}
	d2, err := t.FilterSatelliteByName(satellites, "skywalker")
	if err != nil {
		return model.Response{}, err
	}
	d3, err := t.FilterSatelliteByName(satellites, "sato")
	if err != nil {
		return model.Response{}, err
	}
	position, err := t.triangulationService.GetLocation(p1, p2, p3, d1.Distance, d2.Distance, d3.Distance)
	if err != nil {
		return model.Response{}, err
	}
	err = t.satelliteRepository.DeleteAll()
	if err != nil {
		return model.Response{}, err
	}
	return model.Response{Message: message, Position: position}, nil
}

func (t FuegoQuasarService) FilterSatelliteByName(satellites []model.Satellites, name string) (model.Satellites, error) {
	for _, satellite := range satellites {
		if satellite.Name == name {
			return satellite, nil
		}
	}
	return model.Satellites{}, errors.New("error satellite not found")
}

func (t FuegoQuasarService) GetSatellitesPoints() (model.Point, model.Point, model.Point, error) {
	kenobiX, err := strconv.ParseFloat(os.Getenv("KENOBI_X"), 64)
	if err != nil {
		return model.Point{}, model.Point{}, model.Point{}, err
	}
	kenobiY, err := strconv.ParseFloat(os.Getenv("KENOBI_Y"), 64)
	if err != nil {
		return model.Point{}, model.Point{}, model.Point{}, err
	}

	skywalkerX, err := strconv.ParseFloat(os.Getenv("SKYWALKER_X"), 64)
	if err != nil {
		return model.Point{}, model.Point{}, model.Point{}, err
	}
	skywalkerY, err := strconv.ParseFloat(os.Getenv("SKYWALKER_Y"), 64)
	if err != nil {
		return model.Point{}, model.Point{}, model.Point{}, err
	}

	satoX, err := strconv.ParseFloat(os.Getenv("SATO_X"), 64)
	if err != nil {
		return model.Point{}, model.Point{}, model.Point{}, err
	}
	satoY, err := strconv.ParseFloat(os.Getenv("SATO_Y"), 64)
	if err != nil {
		return model.Point{}, model.Point{}, model.Point{}, err
	}

	return model.Point{X: kenobiX, Y: kenobiY}, model.Point{X: skywalkerX, Y: skywalkerY}, model.Point{X: satoX, Y: satoY}, nil
}

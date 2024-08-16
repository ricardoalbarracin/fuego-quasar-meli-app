package service

import (
	"errors"
	"fmt"
	"fuego-quasar-app/internal/core/domain/port"
	"strings"
)

var (
	ErrZeroMessageLength  = errors.New("message length is zero")
	ErrEmptyMessageResult = errors.New("resulting message is empty")
)

type DecodeMessageService struct {
	logService port.LogService
}

func NewDecodeMessageService(logService port.LogService) port.DecodeMessageService {

	return DecodeMessageService{logService: logService}
}

func (mp DecodeMessageService) GetMessage(message [][]string) (string, error) {
	mp.logService.Info("GetMessage", "message", message)
	messageSize := getMessageLength(message)
	if messageSize == 0 {
		mp.logService.Error("message length is zero")
		return "", fmt.Errorf("message length is zero")
	}

	_, messageLength := getMessageLengthFirtsWord(message, messageSize)
	message = deleteOffset(message, messageLength)

	messageCandidateWords := make([]string, messageSize)
	for index := 0; index < messageSize; index++ {
		messageCandidateWords[index] = getWordByPosition(message, index)
	}

	result := strings.Join(messageCandidateWords, " ")
	if strings.TrimSpace(result) == "" {
		mp.logService.Error("resulting message is empty")
		return "", fmt.Errorf("resulting message is empty")
	}

	return strings.TrimSpace(result), nil
}

func getMessageLength(message [][]string) int {
	messageSize := 0
	for _, msg := range message {
		if len(msg) > messageSize {
			messageSize = len(msg)
		}
	}
	return messageSize
}

func getWordByPosition(message [][]string, index int) string {
	wordCount := make(map[string]int)
	var mostFrequentWord string
	maxCount := 0

	for _, words1 := range message {
		words := words1
		if index < len(words) && words[index] != "" {
			word := words[index]
			wordCount[word]++
			if wordCount[word] > maxCount {
				mostFrequentWord = word
				maxCount = wordCount[word]
			}
		}
	}
	return mostFrequentWord
}

func deleteOffset(message [][]string, mesaggeLength int) [][]string {
	var result [][]string
	for _, words := range message {
		len := len(words)

		offset := len - mesaggeLength
		rangew := words[offset:len]
		result = append(result, rangew)
	}
	return result
}
func getMessageLengthFirtsWord(message [][]string, maxLength int) (int, int) {

	wordCount := make(map[string]int)
	var mostFrequentWord string
	maxCount := 0

	indexMostFrequentWord := 0
	for index := 0; index < maxLength; index++ {
		for indexMessage, words1 := range message {
			words := words1
			if index < len(words) && words[index] != "" {
				word := words[index]
				wordCount[word]++
				if wordCount[word] > maxCount {
					mostFrequentWord = word
					maxCount = wordCount[word]
					indexMostFrequentWord = indexMessage
				}
			}
		}
		if mostFrequentWord != "" {
			return indexMostFrequentWord, len(message[indexMostFrequentWord])
		}
	}

	return 0, 0
}

func removeEmptyStrings(list []string) []string {
	var result []string
	for _, str := range list {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

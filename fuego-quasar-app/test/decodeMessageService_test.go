package service

import (
	"errors"
	"fuego-quasar-app/internal/core/application/service"
	"testing"
)

var (
	ErrZeroMessageLength  = errors.New("message length is zero")
	ErrEmptyMessageResult = errors.New("resulting message is empty")
)

func TestDecodeMessageService_GetMessage(t *testing.T) {
	service := service.NewDecodeMessageService() // Asegúrate de crear la instancia del servicio

	tests := []struct {
		name    string
		message [][]string
		want    string
		wantErr error
	}{
		{
			name: "Simple message with empty strings",
			message: [][]string{
				{"", "este", "es", "un", "mensaje"},
				{"este", "", "un", "mensaje"},
				{"", "este", "es", "", ""}},
			want:    "este es un mensaje",
			wantErr: nil,
		},
		{
			name: "Simple message with empty strings and offset",
			message: [][]string{
				{"este", "", "", "mensaje", ""},
				{"", "es", "", "", "secreto"},
				{"este", "", "un", "", ""}},
			want:    "este es un mensaje secreto",
			wantErr: nil,
		},
		{
			name: "Simple message with empty strings and offset",
			message: [][]string{
				{"", "este", "", "", "mensaje", ""},
				{"", "", "es", "", "", "secreto"},
				{"", "este", "", "un", "", ""}},
			want:    "este es un mensaje secreto",
			wantErr: nil,
		},
		{
			name:    "Simple message with empty strings",
			message: [][]string{{"", "mundo"}, {"hola", ""}, {"hola", "mundo"}},
			want:    "hola mundo",
			wantErr: nil,
		},
		{
			name:    "Simple message with no empty strings",
			message: [][]string{{"hola", "mundo"}, {"hola", "mundo"}, {"hola", "mundo"}},
			want:    "hola mundo",
			wantErr: nil,
		},

		{
			name:    "Empty message",
			message: [][]string{},
			want:    "",
			wantErr: ErrZeroMessageLength,
		},
		{
			name:    "All empty strings",
			message: [][]string{{"", "", ""}, {"", "", ""}, {"", "", ""}},
			want:    "",
			wantErr: ErrEmptyMessageResult,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GetMessage(tt.message)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("GetMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

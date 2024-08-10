package port

type DecodeMessageService interface {
	GetMessage(messages [][]string) (string, error)
}

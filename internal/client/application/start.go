package application

type StartHandler struct{}

func NewStartHandler() *StartHandler { return &StartHandler{} }

func (h *StartHandler) Handle() error {
	return nil
}
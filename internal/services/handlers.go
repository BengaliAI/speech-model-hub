package services

type InferenceRequestHandler struct{}

func (infer *InferenceRequestHandler) SendRequest() (string, error) {
	return "", nil
}

func NewInferenceRequestHandler() *InferenceRequestHandler {
	return &InferenceRequestHandler{}
}

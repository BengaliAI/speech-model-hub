package etypes

import "fmt"

type EType struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e EType) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func (e EType) SetMessage(message string) EType {
	e.Message = message
	return e
}

var ERROR_FILE_FORMAT_NOT_SUPPORTED = EType{
	Code:    1001,
	Message: "File format not supported",
}

var ERROR_INFERENCE_FAILED = EType{
	Code:    1002,
	Message: "Inference failed",
}

var ERROR_MODEL_NOT_FOUND = EType{
	Code:    1003,
	Message: "Model not found",
}

package services

import (
	"mime/multipart"
	"speech-model-hub/internal/domains"
	"speech-model-hub/internal/utils"
)

type ServiceHandler struct{}

func (infer *ServiceHandler) SendRequest(aiModel domains.AIModelInfo, file *multipart.FileHeader) (string, error) {
	type T struct {
		File *multipart.FileHeader `form:"file"`
	}
	body := T{
		File: file,
	}
	_, err := utils.SendPOSTRequest[T](aiModel.URL, aiModel.Authorization, "multipart/form-data", body)
	if err != nil {
		return "", err
	}
	return "", nil
}

func (infer *ServiceHandler) SaveFile(file *multipart.FileHeader) (string, error) {
	name, err := utils.SaveFile("./tmp/data", file)
	if err != nil {
		return "", err
	}
	return name, nil
}

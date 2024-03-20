package services

import (
	"bytes"
	"io"
	"mime/multipart"
	"path/filepath"
	"speech-model-hub/internal/domains"
	"speech-model-hub/internal/utils"
)

type ServiceHandler struct{}

func (infer *ServiceHandler) SendRequest(aiModel domains.AIModelInfo, fileHeader *multipart.FileHeader) (string, error) {
	type T struct {
		Transcript string `json:"transcript"`
	}
	file, err := fileHeader.Open()

	if err != nil {
		return "", err
	}

	defer file.Close()
	body := &bytes.Buffer{}
	formWriter := multipart.NewWriter(body)
	filePart, err := formWriter.CreateFormFile("file", filepath.Base(fileHeader.Filename))
	if err != nil {
		return "", err
	}
	_, err = io.Copy(filePart, file)
	if err != nil {
		return "", err
	}
	err = formWriter.Close()
	if err != nil {
		return "", err
	}
	ret, err := utils.SendPOSTRequestForm[T](aiModel.URL, aiModel.Authorization, formWriter.FormDataContentType(), body)
	if err != nil {
		return "", err
	}
	return ret.Transcript, nil
}

func (infer *ServiceHandler) SaveFile(file *multipart.FileHeader) (string, error) {
	name, err := utils.SaveFile("./tmp/data", file)
	if err != nil {
		return "", err
	}
	return name, nil
}

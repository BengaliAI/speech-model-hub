package services

import (
	"mime/multipart"
	"speech-model-hub/internal/domains"
)

type IFServices interface {
	SendRequest(domains.AIModelInfo, *multipart.FileHeader) (string, error)
	SaveFile(*multipart.FileHeader) (string, error)
}

package domains

import "mime/multipart"

type ModelRepositoryInterface interface {
	GetModelList() ([]AIModelInfo, error)
	GetModelByName(name string) (AIModelInfo, error)
	GetModelByArchitecture(architecture string) (AIModelInfo, error)
	GetModelByAuthor(author string) (AIModelInfo, error)
	AddAIModel(model AIModelInfo) (AIModelInfo, error)
}

type InferenceInterface interface {
	GetInference(file *multipart.FileHeader, display_name string) (string, error)
	SumbitRequest(url string, file *multipart.FileHeader, authorization string) (string, error)
}

type RequestHandlerInterface interface {
	SendRequest() (string, error)
}

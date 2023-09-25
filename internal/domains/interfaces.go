package domains

import "mime/multipart"

type IFModelRepository interface {
	GetModelList() ([]AIModelInfo, error)
	GetModelByDisplayName(name string) (AIModelInfo, error)
	GetModelByName(name string) (AIModelInfo, error)
	GetModelByArchitecture(architecture string) (AIModelInfo, error)
	GetModelByAuthor(author string) (AIModelInfo, error)
	AddAIModel(model AIModelInfo) (AIModelInfo, error)
}

type IFInference interface {
	GetInference(file *multipart.FileHeader, display_name string) (string, error)
	SumbitRequest(url string, file *multipart.FileHeader, authorization string) (string, error)
}

type IFRequestHandler interface {
	SendRequest() (string, error)
}

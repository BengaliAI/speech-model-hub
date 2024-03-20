package domains

type IFModelRepository interface {
	GetModelList() ([]AIModelInfo, error)
	GetModelByDisplayName(name string) (AIModelInfo, error)
	GetModelByName(name string) (AIModelInfo, error)
	GetModelByArchitecture(architecture string) (AIModelInfo, error)
	GetModelByAuthor(author string) (AIModelInfo, error)
	AddAIModel(model AIModelInfo) (AIModelInfo, error)
}

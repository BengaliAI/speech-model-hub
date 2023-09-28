package domains

type DummyDB struct{}

func (db *DummyDB) GetModelList() ([]AIModelInfo, error) {
	return []AIModelInfo{}, nil
}

func (db *DummyDB) GetModelByDisplayName(name string) (AIModelInfo, error) {
	return AIModelInfo{}, nil
}

func (db *DummyDB) GetModelByName(name string) (AIModelInfo, error) {
	return AIModelInfo{}, nil
}

func (db *DummyDB) GetModelByArchitecture(architecture string) (AIModelInfo, error) {
	return AIModelInfo{}, nil
}

func (db *DummyDB) GetModelByAuthor(author string) (AIModelInfo, error) {
	return AIModelInfo{}, nil
}

func (db *DummyDB) AddAIModel(model AIModelInfo) (AIModelInfo, error) {
	return model, nil
}

type DummyRequestHandler struct{}

func (infer *DummyRequestHandler) SendRequest() (string, error) {
	return "", nil
}

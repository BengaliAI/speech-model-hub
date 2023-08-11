package domains

import "mime/multipart"

type AIModelInfo struct {
	Architecture  string `json:"architecture" bson:"architecture"`
	URL           string `json:"url" bson:"url"`
	Authorization string `json:"authorization" bson:"authorization"`
	Author        string `json:"author" bson:"author"`
	DisplayName   string `json:"display_name" bson:"display_name"`
	UpdatedAt     string `json:"updated_at" bson:"updated_at"`
	CreatedAt     string `json:"created_at" bson:"created_at"`
}

type AIInference struct {
	DB            ModelRepositoryInterface
	URL           string
	Authorization string
	Request       RequestHandlerInterface
}

func (infObj *AIInference) sumbitRequest(url string, file *multipart.FileHeader, authorization string) (string, error) {
	return "", nil
}

func (infObj *AIInference) GetInference(file *multipart.FileHeader, display_name string) (string, error) {
	model, err := infObj.DB.GetModelByName(display_name)
	if err != nil {
		return "", err
	}
	return infObj.sumbitRequest(model.URL, file, model.Authorization)
}

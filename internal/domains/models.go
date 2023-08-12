package domains

import (
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AIModelInfo struct {
	Name          string              `json:"name" bson:"name" validate:"required"`
	Architecture  string              `json:"architecture" bson:"architecture" validate:"required"`
	URL           string              `json:"url" bson:"url" validate:"required,url"`
	Authorization string              `json:"authorization" bson:"authorization" validate:"required"`
	Author        string              `json:"author" bson:"author" validate:"required"`
	DisplayName   string              `json:"display_name" bson:"display_name" validate:"required"`
	Active        bool                `json:"active" bson:"active" validate:"required"`
	UpdatedAt     primitive.Timestamp `json:"updated_at" bson:"updated_at" validate:"required"`
	CreatedAt     primitive.Timestamp `json:"created_at" bson:"created_at" validate:"required"`
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
	model, err := infObj.DB.GetModelByDisplayName(display_name)
	if err != nil {
		return "", err
	}
	return infObj.sumbitRequest(model.URL, file, model.Authorization)
}

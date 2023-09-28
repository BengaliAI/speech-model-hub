package domains

import (
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

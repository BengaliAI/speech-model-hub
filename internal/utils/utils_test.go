package utils

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestStruct struct {
	Architecture  string              `json:"architecture" bson:"architecture" validate:"required"`
	URL           string              `json:"url" bson:"url" validate:"required,url"`
	Authorization string              `json:"authorization" bson:"authorization" validate:"required"`
	Author        string              `json:"author" bson:"author" validate:"required"`
	DisplayName   string              `json:"display_name" bson:"display_name" validate:"required"`
	Active        bool                `json:"active" bson:"active" validate:"required"`
	UpdatedAt     primitive.Timestamp `json:"updated_at" bson:"updated_at" validate:"required"`
	CreatedAt     primitive.Timestamp `json:"created_at" bson:"created_at" validate:"required"`
}

func TestValidateStruct(t *testing.T) {
	type args struct {
		model interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case-1",
			args:    args{model: TestStruct{}},
			wantErr: true,
		},
		{
			name: "case-2",
			args: args{
				model: TestStruct{
					Architecture:  "test",
					URL:           "http://google.com",
					Authorization: "test",
					Author:        "test",
					DisplayName:   "test",
					Active:        true,
					UpdatedAt:     primitive.Timestamp{T: 0, I: 0},
					CreatedAt:     primitive.Timestamp{T: 0, I: 0},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateStruct(tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("ValidateStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package infrastructure

import (
	"reflect"
	"speech-model-hub/internal/domains"
	"speech-model-hub/pkg/loggerFx"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestModelRepository_GetModelList(t *testing.T) {
	type fields struct {
		db *MongoDB
	}
	params := DBParams{
		MongoURL:     "mongodb://localhost:27017",
		DatabaseName: "test",
	}
	db := NewDBInstance(params)
	loggerFx.TestSetup()
	tests := []struct {
		name    string
		fields  fields
		want    []domains.AIModelInfo
		wantErr bool
	}{
		{
			name: "test_get_model_list",
			fields: fields{
				db: db,
			},
			want: []domains.AIModelInfo{
				{
					Name:          "test",
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
			repo := &ModelRepository{
				db: tt.fields.db,
			}
			got, err := repo.GetModelList()
			for i := range got {
				got[i].UpdatedAt = primitive.Timestamp{T: 0, I: 0}
				got[i].CreatedAt = primitive.Timestamp{T: 0, I: 0}
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelRepository.GetModelList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ModelRepository.GetModelList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModelRepository_AddAIModel(t *testing.T) {
	type fields struct {
		db *MongoDB
	}
	type args struct {
		model domains.AIModelInfo
	}
	params := DBParams{
		MongoURL:     "mongodb://localhost:27017",
		DatabaseName: "test",
	}
	db := NewDBInstance(params)

	loggerFx.TestSetup()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domains.AIModelInfo
		wantErr bool
	}{
		{
			name: "test_add_ai_model",
			fields: fields{
				db: db,
			},
			args: args{
				model: domains.AIModelInfo{
					Name:          "test",
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
			want: domains.AIModelInfo{
				Name:          "test",
				Architecture:  "test",
				URL:           "http://google.com",
				Authorization: "test",
				Author:        "test",
				DisplayName:   "test",
				Active:        true,
				UpdatedAt:     primitive.Timestamp{T: 0, I: 0},
				CreatedAt:     primitive.Timestamp{T: 0, I: 0},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &ModelRepository{
				db: tt.fields.db,
			}
			got, err := repo.AddAIModel(tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelRepository.AddAIModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ModelRepository.AddAIModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModelRepository_GetModelByDisplayName(t *testing.T) {
	type fields struct {
		db *MongoDB
	}
	type args struct {
		name string
	}
	params := DBParams{
		MongoURL:     "mongodb://localhost:27017",
		DatabaseName: "test",
	}
	db := NewDBInstance(params)
	loggerFx.TestSetup()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domains.AIModelInfo
		wantErr bool
	}{
		{
			name: "test_get_model_by_display_name",
			fields: fields{
				db: db,
			},
			args: args{
				name: "test",
			},
			want: domains.AIModelInfo{
				Name:          "test",
				Architecture:  "test",
				URL:           "http://google.com",
				Authorization: "test",
				Author:        "test",
				DisplayName:   "test",
				Active:        true,
				UpdatedAt:     primitive.Timestamp{T: 0, I: 0},
				CreatedAt:     primitive.Timestamp{T: 0, I: 0},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &ModelRepository{
				db: tt.fields.db,
			}
			got, err := repo.GetModelByDisplayName(tt.args.name)
			got.UpdatedAt = primitive.Timestamp{T: 0, I: 0}
			got.CreatedAt = primitive.Timestamp{T: 0, I: 0}
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelRepository.GetModelByDisplayName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ModelRepository.GetModelByDisplayName() = %v, want %v", got, tt.want)
			}
		})
	}
}

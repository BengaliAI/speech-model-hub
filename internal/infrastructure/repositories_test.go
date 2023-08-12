package infrastructure

import (
	"reflect"
	"speech-model-hub/internal/domains"
	"testing"
)

func TestModelRepository_GetModelList(t *testing.T) {
	type fields struct {
		db *MongoDB
	}
	db, _ := GetDBInstance("mongodb://localhost:27017", "test")
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
			want:    []domains.AIModelInfo{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &ModelRepository{
				db: tt.fields.db,
			}
			got, err := repo.GetModelList()
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
	db, _ := GetDBInstance("mongodb://localhost:27017", "test")
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
				model: domains.AIModelInfo{},
			},
			want:    domains.AIModelInfo{},
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

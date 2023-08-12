package domains

import (
	"mime/multipart"
	"testing"
)

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

func TestInference_GetInference(t *testing.T) {
	type fields struct {
		DB            ModelRepositoryInterface
		URL           string
		Authorization string
		Request       RequestHandlerInterface
	}
	type args struct {
		file         *multipart.FileHeader
		display_name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test_get_inference",
			fields: fields{
				DB:            &DummyDB{},
				URL:           "",
				Authorization: "",
				Request:       &DummyRequestHandler{},
			},
			args: args{
				file:         nil,
				display_name: "",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			infObj := &AIInference{
				DB:            tt.fields.DB,
				URL:           tt.fields.URL,
				Authorization: tt.fields.Authorization,
				Request:       tt.fields.Request,
			}
			t.Log("Inference.GetInference()", infObj)
			got, err := infObj.GetInference(tt.args.file, tt.args.display_name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Inference.GetInference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Inference.GetInference() = %v, want %v", got, tt.want)
			}
		})
	}
}

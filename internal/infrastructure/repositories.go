package infrastructure

import (
	"speech-model-hub/internal/domains"

	"go.mongodb.org/mongo-driver/bson"
)

type ModelRepository struct {
	db *MongoDB
}

func (repo *ModelRepository) GetModelList() ([]domains.AIModelInfo, error) {
	result, err := repo.db.FindManyFromDB("models", bson.D{})
	if err != nil {
		return []domains.AIModelInfo{}, err
	}
	var models []domains.AIModelInfo
	for _, model := range result {
		var temp domains.AIModelInfo
		err = BSONMToStruct(model, &temp)
		if err != nil {
			return []domains.AIModelInfo{}, err
		}
		models = append(models, temp)
	}
	return models, nil
}

func (repo *ModelRepository) GetModelByName(name string) (domains.AIModelInfo, error) {
	result, err := repo.db.FindOneFromDB("models", bson.D{{Key: "display_name", Value: name}})
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	var model domains.AIModelInfo
	err = BSONMToStruct(result, &model)
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	return model, nil
}

func (repo *ModelRepository) GetModelByArchitecture(architecture string) (domains.AIModelInfo, error) {
	result, err := repo.db.FindOneFromDB("models", bson.D{{Key: "architecture", Value: architecture}})
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	var model domains.AIModelInfo
	err = BSONMToStruct(result, &model)
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	return model, nil
}

func (repo *ModelRepository) GetModelByAuthor(author string) (domains.AIModelInfo, error) {
	result, err := repo.db.FindOneFromDB("models", bson.D{{Key: "author", Value: author}})
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	var model domains.AIModelInfo
	err = BSONMToStruct(result, &model)
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	return model, nil
}

func (repo *ModelRepository) AddAIModel(model domains.AIModelInfo) (domains.AIModelInfo, error) {
	err := repo.db.InsertIntoDB("models", model)
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	return model, nil
}

func GetModelRepository() domains.ModelRepositoryInterface {
	db, err := GetDBInstance()
	if err != nil {
		panic(err)
	}
	return &ModelRepository{
		db: db,
	}
}

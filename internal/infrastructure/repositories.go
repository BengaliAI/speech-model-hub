package infrastructure

import (
	"speech-model-hub/internal/domains"
	"speech-model-hub/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type ModelRepository struct {
	db *MongoDB
}

func (repo *ModelRepository) GetModelList() ([]domains.AIModelInfo, error) {
	result, err := repo.db.FindManyFromDB("models", bson.D{{Key: "active", Value: true}})
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
		err = utils.ValidateStruct(temp)
		if err != nil {
			return []domains.AIModelInfo{}, err
		}
		models = append(models, temp)
	}
	return models, nil
}

func (repo *ModelRepository) GetModelByDisplayName(name string) (domains.AIModelInfo, error) {
	result, err := repo.db.FindOneFromDB("models", bson.D{{Key: "display_name", Value: name}})
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	var model domains.AIModelInfo
	err = BSONMToStruct(result, &model)
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	err = utils.ValidateStruct(model)
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
	err = utils.ValidateStruct(model)
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
	err = utils.ValidateStruct(model)
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	return model, nil
}

func (repo *ModelRepository) GetModelByName(name string) (domains.AIModelInfo, error) {
	result, err := repo.db.FindOneFromDB("models", bson.D{{Key: "name", Value: name}})
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	var model domains.AIModelInfo
	err = BSONMToStruct(result, &model)
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	err = utils.ValidateStruct(model)
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	return model, nil
}

func (repo *ModelRepository) AddAIModel(model domains.AIModelInfo) (domains.AIModelInfo, error) {
	err := utils.ValidateStruct(model)
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	err = repo.db.InsertIntoDB("models", model)
	if err != nil {
		return domains.AIModelInfo{}, err
	}
	return model, nil
}

func GetModelRepository(mongoURL string, databaseName string) (domains.ModelRepositoryInterface, error) {
	db, err := GetDBInstance(mongoURL, databaseName)
	if err != nil {
		return nil, err
	}
	return &ModelRepository{
		db: db,
	}, nil
}

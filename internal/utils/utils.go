package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetRandomString(length int) string {
	return stringWithCharset(length, charset)
}

func GenerateUUID4() string {
	return uuid.NewString()
}

func ValidateStruct(model interface{}) error {
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(path string, file *multipart.FileHeader) (string, error) {
	err := CreateDirIfNotExist(path)
	if err != nil {
		return "", err
	}
	newFileName := "./" + filepath.Join(path, GetRandomString(12)+"-"+file.Filename)
	newFile, err := os.Create(newFileName)
	if err != nil {
		return "", err
	}
	defer newFile.Close()
	oldFile, err := file.Open()
	if err != nil {
		return "", err
	}
	_, err = io.Copy(newFile, oldFile)
	if err != nil {
		return "", err
	}
	return newFileName, nil
}

func CreateDirIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func SendPOSTRequest[T interface{}](url, authToken, contentType string, body interface{}) (T, error) {
	var res T
	payload, err := json.Marshal(body)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

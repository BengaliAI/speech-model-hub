package infrastructure

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setUniqueIndex(db *MongoDB, collectionName string, uniques bson.D) (bool, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	collection := db.GetCollection(collectionName)
	if len(uniques) != 0 {
		collection.Indexes().CreateOne(ctx, mongo.IndexModel{Keys: uniques, Options: options.Index().SetUnique(true)})
	}
	return true, nil
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := os.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("failed parsing pem file")
	}

	return tlsConfig, nil
}

func BSONMToStruct(data bson.M, result interface{}) error {
	bytes, err := bson.Marshal(data)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(bytes, result)
	if err != nil {
		return err
	}
	return nil
}

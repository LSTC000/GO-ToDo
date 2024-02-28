package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient(ctx context.Context, host, port, username, password, database string) (*mongo.Database, error) {
	var isAuth bool
	var mongoURI string

	if username != "" && password != "" {
		isAuth = true
		mongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s", host, port, username, password)
	} else {
		mongoURI = fmt.Sprintf("mongodb://%s:%s", host, port)
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	if isAuth {
		clientOptions.SetAuth(options.Credential{
			Username: username,
			Password: password,
		})
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("connect error: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping error: %w", err)
	}

	return client.Database(database), nil
}

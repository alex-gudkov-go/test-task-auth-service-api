package mongo_store

import (
	"context"
	"test-task-auth-service-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore struct {
	db *mongo.Database
}

type RefreshTokenEntity struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	UserId string             `bson:"userId"`
	Value  string             `bson:"value"`
}

func New(db *mongo.Database) *MongoStore {
	return &MongoStore{db}
}

func createRefreshTokenEntity(refreshToken *models.RefreshToken) *RefreshTokenEntity {
	return &RefreshTokenEntity{
		// Id: will be generated by MongoDB
		UserId: refreshToken.UserId,
		Value:  refreshToken.Value,
	}
}

func toRefreshTokenModel(refreshTokenEntity *RefreshTokenEntity) *models.RefreshToken {
	return &models.RefreshToken{
		Id:     refreshTokenEntity.Id.Hex(),
		UserId: refreshTokenEntity.UserId,
		Value:  refreshTokenEntity.Value,
	}
}

func (ms *MongoStore) SaveRefreshToken(ctx context.Context, refreshToken *models.RefreshToken) error {
	refreshTokenEntity := createRefreshTokenEntity(refreshToken)

	res, err := ms.db.Collection("refresh_tokens").InsertOne(ctx, refreshTokenEntity)
	if err != nil {
		return err
	}

	// set ID from MongoDB
	refreshToken.Id = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (ms *MongoStore) FindRefreshTokenById(ctx context.Context, id string) (*models.RefreshToken, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	refreshTokenEntity := new(RefreshTokenEntity)
	err = ms.db.Collection("refresh_tokens").FindOne(ctx, bson.M{"_id": objectId}).Decode(refreshTokenEntity)
	if err != nil {
		return nil, err
	}

	return toRefreshTokenModel(refreshTokenEntity), nil
}

func (ms *MongoStore) DeleteRefreshTokenById(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = ms.db.Collection("refresh_tokens").DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	return nil
}

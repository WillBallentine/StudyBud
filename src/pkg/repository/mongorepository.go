package repository

import (
	"context"

	"studybud/src/cmd/utils"
	"studybud/src/pkg/entity"

	//"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository interface {
	AddUser(user entity.User, ctx context.Context) (primitive.ObjectID, error)
	GetUserById(oId primitive.ObjectID, ctx context.Context) (*entity.User, error)
	GetUserByEmail(email string, ctx context.Context) (*entity.User, bool)
}

type mongoRepository struct {
	client *mongo.Client
	config *utils.Configuration
}

func NewMongoRepository(config *utils.Configuration, client *mongo.Client) MongoRepository {
	return &mongoRepository{config: config, client: client}
}

func (userRepo *mongoRepository) AddUser(user entity.User, ctx context.Context) (primitive.ObjectID, error) {
	collection := userRepo.client.Database(userRepo.config.Database.DbName).Collection(userRepo.config.Database.Collection)

	insertResult, err := collection.InsertOne(ctx, user)

	if err != mongo.ErrNilCursor {
		return primitive.NilObjectID, err
	}

	if oidResult, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		return oidResult, nil
	} else {
		return primitive.NilObjectID, err
	}
}

func (userRepo *mongoRepository) GetUserById(oId primitive.ObjectID, ctx context.Context) (*entity.User, error) {

	collection := userRepo.client.Database(userRepo.config.Database.DbName).Collection(userRepo.config.Database.Collection)

	filter := bson.D{primitive.E{Key: "_id", Value: oId}}

	var user *entity.User

	collection.FindOne(ctx, filter).Decode(&user)

	return user, nil
}

func (userRepo *mongoRepository) GetUserByEmail(email string, ctx context.Context) (*entity.User, bool) {
	collection := userRepo.client.Database(userRepo.config.Database.DbName).Collection(userRepo.config.Database.Collection)

	filter := bson.D{primitive.E{Key: "email", Value: email}}
	var exists bool

	var user *entity.User

	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		exists = false
	} else {
		exists = true
	}

	return user, exists
}

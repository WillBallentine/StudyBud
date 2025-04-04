package repository

import (
	"context"

	"studybud/src/cmd/utils"
	"studybud/src/pkg/entity"
	model "studybud/src/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository interface {
	AddUser(user entity.User, ctx context.Context) (primitive.ObjectID, error)
	UpsertSessionInfo(id primitive.ObjectID, st string, ct string, ctx context.Context) (bool, error)
	GetUserById(oId primitive.ObjectID, ctx context.Context) (*entity.User, error)
	GetUserByEmail(email string, ctx context.Context) (*entity.User, bool)
	AddSyllabus(syll entity.SyllabusDataEntity, ctx context.Context) (primitive.ObjectID, error)
	AddStudyPlan(plan model.StudyPlan, ctx context.Context) (primitive.ObjectID, error)
	GetStudyPlanByID(oId primitive.ObjectID, ctx context.Context) (*model.StudyPlan, error)
	UpsertSyllabusId(userId primitive.ObjectID, syllId primitive.ObjectID, ctx context.Context) (bool, error)
	UpsertStudyPlanId(userId primitive.ObjectID, planId primitive.ObjectID, ctx context.Context) (bool, error)
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

func (userRepo *mongoRepository) UpsertSessionInfo(id primitive.ObjectID, st string, ct string, ctx context.Context) (bool, error) {
	collection := userRepo.client.Database(userRepo.config.Database.DbName).Collection(userRepo.config.Database.Collection)
	var updated bool

	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	update := bson.M{
		"$set": bson.M{
			"session_token": st,
			"csrf_token":    ct,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		updated = false
		return updated, err
	} else {
		updated = true
	}

	return updated, nil

}

func (userRepo *mongoRepository) UpsertSyllabusId(userId primitive.ObjectID, syllId primitive.ObjectID, ctx context.Context) (bool, error) {
	collection := userRepo.client.Database(userRepo.config.Database.DbName).Collection(userRepo.config.Database.Collection)

	filter := bson.D{primitive.E{Key: "_id", Value: userId}}
	update := bson.M{
		"$push": bson.M{"syllabi": syllId},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"syllabi": bson.A{syllId}}})
		if err != nil {
			return false, err
		}
	}

	return true, nil

}

func (userRepo *mongoRepository) UpsertStudyPlanId(userId primitive.ObjectID, planId primitive.ObjectID, ctx context.Context) (bool, error) {
	collection := userRepo.client.Database(userRepo.config.Database.DbName).Collection(userRepo.config.Database.Collection)

	filter := bson.D{primitive.E{Key: "_id", Value: userId}}
	update := bson.M{
		"$push": bson.M{"study_plans": planId},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"study_plans": bson.A{planId}}})
		if err != nil {
			return false, err
		}
	}

	return true, nil

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

func (syllRepo *mongoRepository) AddSyllabus(syll entity.SyllabusDataEntity, ctx context.Context) (primitive.ObjectID, error) {
	collection := syllRepo.client.Database(syllRepo.config.Database.DbName).Collection(syllRepo.config.Database.Collection)

	insertResult, err := collection.InsertOne(ctx, syll)

	if err != nil {
		return primitive.NilObjectID, err
	}

	if oidResult, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		return oidResult, nil
	} else {
		return primitive.NilObjectID, err
	}
}

func (planRepo *mongoRepository) AddStudyPlan(plan model.StudyPlan, ctx context.Context) (primitive.ObjectID, error) {
	collection := planRepo.client.Database(planRepo.config.Database.DbName).Collection(planRepo.config.Database.Collection)

	insertResult, err := collection.InsertOne(ctx, plan)

	if err != nil {
		return primitive.NilObjectID, err
	}

	if oidResult, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		return oidResult, nil
	} else {
		return primitive.NilObjectID, err
	}
}

func (planRepo *mongoRepository) GetStudyPlanByID(oId primitive.ObjectID, ctx context.Context) (*model.StudyPlan, error) {

	collection := planRepo.client.Database(planRepo.config.Database.DbName).Collection(planRepo.config.Database.Collection)

	filter := bson.D{primitive.E{Key: "_id", Value: oId}}

	var plan *model.StudyPlan

	collection.FindOne(ctx, filter).Decode(&plan)

	return plan, nil
}

package repositoryimpl

import (
	"auth/modules/user/data/model"
	"auth/modules/user/domain/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	collection *mongo.Collection
}


func NewUserRepository(collection *mongo.Collection) repository.UserRepository {
	return &UserRepositoryImpl{
		collection: collection,
	}
}

func (repo *UserRepositoryImpl)  CreateUser(ctx context.Context,user *model.User)( *model.User,error) {
	res,err :=repo.collection.InsertOne(ctx,user)
	if err != nil {
		return nil,err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}
	user.ID = oid
	return user, nil
}

func (repo *UserRepositoryImpl)  UpdateUser(ctx context.Context,user *model.User)( *model.User,error) {
	res,err :=repo.collection.UpdateOne(ctx,bson.M{"_id": user.ID},bson.M{"$set": user})
	if err != nil {
		return nil,err
	}
	if res.ModifiedCount == 0 {
		return nil,nil
	}
	return user,nil
}

func (repo *UserRepositoryImpl) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	cur, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user model.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users,user)
	}
	return users, nil
}
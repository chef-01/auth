package repository

import (
	"auth/modules/user/data/model"
	"context"
)


type UserRepository interface {
 CreateUser(ctx context.Context, user *model.User)( *model.User,error)
 UpdateUser(ctx context.Context, user *model.User)( *model.User,error)
 GetAllUsers(ctx context.Context)( []model.User,error)
 

}
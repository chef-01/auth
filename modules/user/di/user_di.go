package di

import (
	"auth/common/config"
	"auth/modules/user/domain/usecases"
	"auth/modules/user/presentation/controllers"
	"auth/modules/user/presentation/router"
	"auth/modules/user/data/repository_impl"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
Execute initializes all dependencies and sets up the router.

Params:
- db: MongoDB client instance.
- config: Application configuration.

Returns:
- *mux.Router: The fully configured router ready to be used by the server.
*/
func Execute(db *mongo.Database, cfg *config.Config) *mux.Router {
	// Initialize services
	otpService := usecases.NewTwilioOTPService()
	jwtService := usecases.NewJWTService(*cfg)

	// Initialize repository
	userCollection := db.Collection("users")
	userRepo := repositoryimpl.NewUserRepository(userCollection)

	// Initialize use cases
	userUseCase := usecases.NewUserUseCase(userRepo, otpService, jwtService)

	// Initialize controllers
	userController := controllers.NewUserController(userUseCase)

	// Initialize router
	r := mux.NewRouter()
	router.SetupUserRoutes(r, userController)

	return r
}

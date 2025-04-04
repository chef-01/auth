package controllers

import (
	"auth/common/helpers"
	"auth/common/strings"
	"auth/modules/user/domain/usecases"
	"auth/modules/user/presentation/dto"
	"net/http"
)

type UserController struct {
	userUseCase *usecases.UserUseCase
}

func NewUserController(uc *usecases.UserUseCase) *UserController {
	return &UserController{userUseCase: uc}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	req, ok := r.Context().Value(dto.CreateUserRequestKey).(dto.CreateUserRequest)
	if !ok {
		helpers.JSONResponse(w, http.StatusInternalServerError, false, strings.RequestValidationFailedMsg, nil, nil)
		return
	}

	userModel, err := req.ToModel()
	if err != nil {
		helpers.JSONResponse(w, http.StatusBadRequest, false, strings.InvalidDataMsg+": "+err.Error(), nil, nil)
		return
	}

	user, err := uc.userUseCase.CreateUser(r.Context(), userModel)
	if err != nil {
		helpers.JSONResponse(w, http.StatusInternalServerError, false, err.Error(), nil, nil)
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, true, "user "+strings.CreateSuccessMsg, nil, user)
}


func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	req, ok := r.Context().Value(dto.UpdateUserRequestKey).(dto.UpdateUserRequest)
	if !ok {
		helpers.JSONResponse(w, http.StatusInternalServerError, false, strings.RequestValidationFailedMsg, nil, nil)
		return
	}

	id := r.Context().Value(dto.CreateUserRequestKey).(string)
	userModel, err := req.ToModel(id)
	if err != nil {
		helpers.JSONResponse(w, http.StatusBadRequest, false, strings.InvalidDataMsg+": "+err.Error(), nil, nil)
		return
	}

	user, err := uc.userUseCase.UpdateUser(r.Context(), userModel)
	if err != nil {
		helpers.JSONResponse(w, http.StatusInternalServerError, false, err.Error(), nil, nil)
		return
	}

	helpers.JSONResponse(w, http.StatusOK, true, "user "+strings.UpdateSuccessMsg, nil, user)
}

func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.userUseCase.GetAllUsers(r.Context())
	if err != nil {
		helpers.JSONResponse(w, http.StatusInternalServerError, false, err.Error(), nil, nil)
		return
	}

	if len(users) == 0 {
		helpers.JSONResponse(w, http.StatusOK, true, "no users found", nil, nil)
		return
	}

	helpers.JSONResponse(w, http.StatusOK, true, "users retrieved successfully", nil, users)
}

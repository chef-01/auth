package dto

import (
	"auth/common/strings"
	"auth/modules/user/data/model"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CreateUserRequestKey   = "createUserRequest"
	UpdateUserRequestKey   = "updateUserRequest"
	LoginRequestKey        = "loginRequest"
	GenerateOTPRequestKey  = "generateOTPRequest"
)

// CreateUserRequest defines the payload for creating a new user.
type CreateUserRequest struct {
	CommunityID  string           `json:"community_id" validate:"required"`
	Phones       []string         `json:"phones" validate:"required,dive,required"`
	Emails       []string         `json:"emails" validate:"required,dive,required,email"`
	RoleID       string           `json:"role_id" validate:"required"`
	Name         string           `json:"name" validate:"required,min=2"`
	PPUrl        string           `json:"pp_url,omitempty"`
	Status       string           `json:"status" validate:"required"`
	Addresses    []AddressRequest `json:"addresses" validate:"required,dive"`
	DepartmentID string           `json:"department_id" validate:"required"`
}

// AddressRequest defines the payload for an address.
type AddressRequest struct {
	Street1 string  `json:"street1" validate:"required"`
	Street2 string  `json:"street2,omitempty"`
	City    string  `json:"city" validate:"required"`
	State   string  `json:"state,omitempty"`
	Country string  `json:"country" validate:"required"`
	Pincode string  `json:"pincode" validate:"required"`
	Lat     float64 `json:"lat" validate:"required"`
	Lng     float64 `json:"lng" validate:"required"`
}

// ToModel converts CreateUserRequest into a model.User.
func (req *CreateUserRequest) ToModel() (*model.User, error) {
	communityOID, err := primitive.ObjectIDFromHex(req.CommunityID)
	if err != nil {
		return nil, errors.New("community " + strings.InvalidIDMsg)
	}
	roleOID, err := primitive.ObjectIDFromHex(req.RoleID)
	if err != nil {
		return nil, errors.New("role " + strings.InvalidIDMsg)
	}
	deptOID, err := primitive.ObjectIDFromHex(req.DepartmentID)
	if err != nil {
		return nil, errors.New("department " + strings.InvalidIDMsg)
	}

	var addresses []model.Address
	for _, addr := range req.Addresses {
		addresses = append(addresses, model.Address{
			Street1: addr.Street1,
			Street2: addr.Street2,
			City:    addr.City,
			State:   addr.State,
			Country: addr.Country,
			Pincode: addr.Pincode,
			Coordinates: model.Coordinates{
				Lat: addr.Lat,
				Lng: addr.Lng,
			},
		})
	}

	// Convert each phone string to a PhoneInfo.
	var phoneInfos []model.PhoneInfo
	for _, p := range req.Phones {
		phoneInfos = append(phoneInfos, model.PhoneInfo{
			Number: p,
		})
	}

	// Convert each email string to an EmailInfo.
	var emailInfos []model.EmailInfo
	for _, e := range req.Emails {
		emailInfos = append(emailInfos, model.EmailInfo{
			Email: e,
		})
	}

	return &model.User{
		CommunityID:  communityOID,
		RoleID:       roleOID,
		Name:         req.Name,
		PPUrl:        req.PPUrl,
		Status:       req.Status,
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
		Addresses:    addresses,
		DepartmentID: deptOID,
		Phones:       phoneInfos,
		Emails:       emailInfos,
	}, nil
}

// UpdateUserRequest defines the payload for updating a user's details.
type UpdateUserRequest struct {
	Email  string `json:"email,omitempty" validate:"omitempty,email"`
	Name   string `json:"name,omitempty" validate:"omitempty,min=2"`
	PPUrl  string `json:"pp_url,omitempty"`
	Status string `json:"status,omitempty"`
}

// ToModel converts UpdateUserRequest into a model.User. The id parameter is used to set the user ID.
func (req *UpdateUserRequest) ToModel(id string) (*model.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("user " + strings.InvalidIDMsg)
	}
	return &model.User{
		ID:         oid,
		Name:       req.Name,
		PPUrl:      req.PPUrl,
		Status:     req.Status,
		ModifiedAt: time.Now(),
	}, nil
}

// LoginRequest defines the payload for a user login.
type LoginRequest struct {
	Phone string `json:"phone" validate:"required"`
	OTP   string `json:"otp" validate:"required"`
}

// GenerateOTPRequest defines the payload for generating an OTP.
type GenerateOTPRequest struct {
	Phone string `json:"phone" validate:"required"`
}

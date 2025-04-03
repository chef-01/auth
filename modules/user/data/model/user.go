package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Coordinates defines the latitude and longitude.
type Coordinates struct {
	Lat float64 `bson:"lat" json:"lat"`
	Lng float64 `bson:"lng" json:"lng"`
}

// Address represents a physical address with detailed fields.
type Address struct {
	Street1     string      `bson:"street1" json:"street1"`
	Street2     string      `bson:"street2,omitempty" json:"street2,omitempty"`
	City        string      `bson:"city" json:"city"`
	State       string      `bson:"state,omitempty" json:"state,omitempty"`
	Country     string      `bson:"country" json:"country"`
	Pincode     string      `bson:"pincode" json:"pincode"`
	Coordinates Coordinates `bson:"coordinates" json:"coordinates"`
}

// PhoneInfo groups phone number and OTP-related fields.
type PhoneInfo struct {
	Number       string    `bson:"number" json:"number"`
	OTP          string    `bson:"otp,omitempty" json:"-"`
	OTPExpiresAt time.Time `bson:"otp_expires_at,omitempty" json:"-"`
	IsPrimary    bool      `bson:"is_primary" json:"is_primary"`
}

// EmailInfo groups email address and OTP-related fields.
type EmailInfo struct {
	Email       string    `bson:"email" json:"email"`
	OTP          string    `bson:"otp,omitempty" json:"-"`
	OTPExpiresAt time.Time `bson:"otp_expires_at,omitempty" json:"-"`
}

// User represents a user in the Primus Web Application.
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CommunityID  primitive.ObjectID `bson:"community_id" json:"community_id"`
	Phones       []PhoneInfo        `bson:"phone_info" json:"phone_info"`
	Emails       []EmailInfo        `bson:"email_info,omitempty" json:"email_info,omitempty"`
	RoleID       primitive.ObjectID `bson:"role_id" json:"role_id"`
	Name         string             `bson:"name" json:"name"`
	PPUrl        string             `bson:"pp_url" json:"pp_url"` // Profile picture URL
	Status       string             `bson:"status" json:"status"` // e.g., active/inactive
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	ModifiedAt   time.Time          `bson:"modified_at" json:"modified_at"`
	Addresses    []Address          `bson:"addresses" json:"addresses"`
	DepartmentID primitive.ObjectID `bson:"department_id" json:"department_id"`
}
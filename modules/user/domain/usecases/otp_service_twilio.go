package usecases

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"math/rand"
)

type TwilioOTPService struct {
	client    *twilio.RestClient
	fromPhone string
}

func NewTwilioOTPService() OTPService {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromPhone := os.Getenv("TWILIO_FROM_PHONE")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	return &TwilioOTPService{
		client:    client,
		fromPhone: fromPhone,
	}
}

func (s *TwilioOTPService) GenerateOtp(phone string) (string, error) {
	otp := strconv.Itoa(rand.Intn(900000) + 100000)
	messageBody := fmt.Sprintf("Your OTP is %s", otp)

	params := &openapi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(s.fromPhone)
	params.SetBody(messageBody)

	_, err := s.client.Api.CreateMessage(params)
	if err != nil {
		return "", fmt.Errorf("failed to send otp: %w", err)
	}

	return otp, nil
}
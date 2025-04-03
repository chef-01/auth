package usecases


type OTPService interface {
	GenerateOtp(phone string) (string, error)
}
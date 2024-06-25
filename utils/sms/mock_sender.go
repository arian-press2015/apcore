package sms

import (
	"fmt"
)

type MockSmsSender struct{}

func NewMockSmsSender() *MockSmsSender {
	return &MockSmsSender{}
}

func (s *MockSmsSender) SendLoginOtp(otp string, phone string) error {
	fmt.Printf("Mock OTP SMS sent to %s: %s\n", phone, otp)
	return nil
}

func (s *MockSmsSender) SendInvoice(invoice string, phone string) error {
	fmt.Printf("Mock Invoice SMS sent to %s: %s\n", phone, invoice)
	return nil
}

func (s *MockSmsSender) SendBulkSms(message string, phones []string) error {
	fmt.Printf("Mock Bulk SMS sent to %v with message: %s\n", phones, message)
	return nil
}

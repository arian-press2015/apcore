package sms

import "apcore/config"

type SmsSender interface {
	SendLoginOtp(otp string, phone string) error
	SendInvoice(invoice string, phone string) error
	SendBulkSms(message string, phones []string) error
}

func NewSmsSender(config *config.Config) SmsSender {
	if config.Env == "development" {
		return NewMockSmsSender()
	} else {
		return NewSmsIrSender(config.Sms.ApiUrl, config.Sms.ApiKey, config.Sms.LineNumber)
	}
}

package sms

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	OTP_TEMPLATE_ID     = 100000
	INVOICE_TEMPLATE_ID = 200000
)

type SmsIrSender struct {
	apiUrl     string
	apiKey     string
	lineNumber string
}

type SmsRequestParams struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SmsRequest struct {
	Mobile     string             `json:"mobile"`
	TemplateID int                `json:"templateId"`
	Parameters []SmsRequestParams `json:"parameters"`
}

type BulkSmsRequest struct {
	Message      string   `json:"message"`
	Phones       []string `json:"phones"`
	SendDateTime *string  `json:"sendDateTime"`
}

func NewSmsIrSender(apiUrl, apiKey, lineNumber string) *SmsIrSender {
	return &SmsIrSender{
		apiUrl:     apiUrl,
		apiKey:     apiKey,
		lineNumber: lineNumber,
	}
}

func (s *SmsIrSender) SendSms(phone string, templateId int, params []SmsRequestParams) error {
	smsRequest := SmsRequest{
		Mobile:     phone,
		TemplateID: templateId,
		Parameters: params,
	}

	jsonData, err := json.Marshal(smsRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.apiUrl+"/send/verify", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("x-api-key", s.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to send SMS")
	}

	return nil
}

func (s *SmsIrSender) SendLoginOtp(otp string, phone string) error {
	fmt.Println("api url:", s.apiUrl, s.apiKey, s.lineNumber)
	params := []SmsRequestParams{
		{Name: "CODE", Value: otp},
	}
	err := s.SendSms(phone, OTP_TEMPLATE_ID, params)
	return err
}

func (s *SmsIrSender) SendInvoice(invoice string, phone string) error {
	params := []SmsRequestParams{
		{Name: "INVOICE", Value: invoice},
	}
	err := s.SendSms(phone, INVOICE_TEMPLATE_ID, params)
	return err
}

func (s *SmsIrSender) SendBulkSms(message string, phones []string) error {
	bulkSmsRequest := BulkSmsRequest{
		Message:      message,
		Phones:       phones,
		SendDateTime: nil,
	}

	jsonData, err := json.Marshal(bulkSmsRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.apiUrl+"/send/bulk", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", s.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to send bulk SMS")
	}

	return nil
}

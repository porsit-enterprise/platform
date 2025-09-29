package sms

import (
	"errors"
	"log/slog"

	cfg_entities "github.com/porsit-enterprise/platform/foundation/configuration/entities"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

type SmsProvider interface {
	SendOTP(receiver string, password string) (Result, error)
	Status(messageId string) (Result, error)
}

type Result struct {
	MessageId string
	Status    MessageStatus
}

var (
	ErrHTTP = errors.New("")
	ErrAPI  = errors.New("")
	ErrMisc = errors.New("")
)

type MessageStatus int

const (
	STATUS_QUEUED       MessageStatus = 1
	STATUS_SCHULDED     MessageStatus = 2
	STATUS_SENTTOCENTER MessageStatus = 3
	STATUS_SENT         MessageStatus = 4
	STATUS_FAILED       MessageStatus = 5
	STATUS_DELIVERED    MessageStatus = 6
	STATUS_UNDELIVERED  MessageStatus = 7
	STATUS_FILTERED     MessageStatus = 8
	STATUS_UNKNOWN      MessageStatus = 9
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func New(cfg cfg_entities.ProviderSMS) SmsProvider {
	slog.Info("connecting to SMS provider")

	provider := newKavenegar(cfg.Kavenegar)
	return provider
}

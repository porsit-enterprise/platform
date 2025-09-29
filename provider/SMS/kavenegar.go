package sms

import (
	"fmt"
	"log/slog"

	"github.com/kavenegar/kavenegar-go"

	cfg_entities "github.com/porsit-enterprise/platform/foundation/configuration/entities"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

type kavenegarSmsProvider struct {
	client *kavenegar.Kavenegar
	cfg    cfg_entities.Kavenegar
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func newKavenegar(cfg cfg_entities.Kavenegar) kavenegarSmsProvider {
	api := kavenegar.New(cfg.ApiKey)
	return kavenegarSmsProvider{
		client: api,
		cfg:    cfg,
	}
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func (k kavenegarSmsProvider) mapStatus(kstatus kavenegar.MessageStatusType) MessageStatus {
	switch kstatus {
	case kavenegar.Type_MessageStatus_Queued:
		return STATUS_QUEUED
	case kavenegar.Type_MessageStatus_Schulded:
		return STATUS_SCHULDED
	case kavenegar.Type_MessageStatus_SentToCenter:
		return STATUS_SENTTOCENTER
	case kavenegar.Type_MessageStatus_Sent:
		return STATUS_SENT
	case kavenegar.Type_MessageStatus_Failed:
		return STATUS_FAILED
	case kavenegar.Type_MessageStatus_Delivered:
		return STATUS_DELIVERED
	case kavenegar.Type_MessageStatus_Undelivered:
		return STATUS_UNDELIVERED
	case kavenegar.Type_MessageStatus_Filtered:
		return STATUS_FILTERED
	}
	return STATUS_UNKNOWN
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func (k kavenegarSmsProvider) SendOTP(receiver string, password string) (Result, error) {
	params := &kavenegar.VerifyLookupParam{Type: kavenegar.Type_VerifyLookup_Sms}
	res, err := k.client.Verify.Lookup(receiver, k.cfg.OtpTemplate, password, params)
	if err != nil {
		slog.Error("error in send otp by kavenegar", slog.Any("err", err))
		switch err.(type) {
		case *kavenegar.APIError:
			return Result{}, ErrAPI
		case *kavenegar.HTTPError:
			return Result{}, ErrHTTP
		default:
			return Result{}, ErrMisc
		}
	}
	return Result{
		MessageId: fmt.Sprintf("%d", res.MessageID),
		Status:    k.mapStatus(res.Status),
	}, nil
}

func (k kavenegarSmsProvider) Status(messageId string) (Result, error) {
	status, err := k.client.Message.Status([]string{messageId})
	if err != nil {
		slog.Error("error in getting status by kavenegar", slog.Any("err", err))
		switch err.(type) {
		case *kavenegar.APIError:
			return Result{}, ErrAPI
		case *kavenegar.HTTPError:
			return Result{}, ErrHTTP
		default:
			return Result{}, ErrMisc
		}
	}
	if len(status) == 0 {
		return Result{}, fmt.Errorf("no data")
	}
	return Result{
		MessageId: fmt.Sprintf("%d", status[0].MessageId),
		Status:    k.mapStatus(kavenegar.MessageStatusType(status[0].Status)),
	}, nil
}

package ed

import (
	"errors"
	"reflect"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridService struct {
	apiTesting bool
	client     *sendgrid.Client
}

func NewSendGridService(apiKey string, testing bool) (*SendGridService, error) {
	if apiKey == "" {
		return nil, errors.New("sendgrid api key is missing")
	}
	client := sendgrid.NewSendClient(apiKey)
	s := &SendGridService{
		apiTesting: testing,
		client:     client,
	}
	return s, nil
}

func (s SendGridService) Send(e Email) (interface{}, error) {

	if reflect.DeepEqual(s, SendGridService{}) {
		return nil, errors.New("SendGridService struct must be properly initialised")
	}

	emailPlainText, err := e.PlainText()
	if err != nil {
		return "", err
	}
	emailHTML, err := e.HTML()
	if err != nil {
		return "", err
	}
	switch "" {
	case e.FromName:
		return "", errors.New("from name not set")
	case e.FromEmail:
		return "", errors.New("from email not set")
	case e.RecipientEmail:
		return "", errors.New("recipient email not set")
	case e.Subject:
		return "", errors.New("subject not set")
	default:
	}

	message := mail.NewSingleEmail(mail.NewEmail(e.FromName, e.FromEmail), e.Subject, mail.NewEmail("", e.RecipientEmail), emailPlainText, emailHTML)

	if s.apiTesting {
		return message, nil
	}

	response, err := s.client.Send(message)
	if err != nil {
		return response, nil
	}

	return response, err
}

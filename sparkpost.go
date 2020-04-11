package ed

import (
	"errors"
	"fmt"
	"reflect"

	sp "github.com/SparkPost/gosparkpost"
)

type SparkPostService struct {
	apiTesting bool
	client     sp.Client
}

func NewSparkPostService(apiKey string, apiEndpoint string, apiTesting bool) (*SparkPostService, error) {
	if apiKey == "" {
		return nil, errors.New("sparkpost api key is missing")
	}

	if apiEndpoint == "" {
		return nil, errors.New("sparkpost api endpoint is missing")
	}

	config := &sp.Config{
		BaseUrl:    apiEndpoint,
		ApiKey:     apiKey,
		ApiVersion: 1,
	}
	var client sp.Client
	err := client.Init(config)
	s := &SparkPostService{
		apiTesting: apiTesting,
		client:     client,
	}
	return s, err
}

func (s SparkPostService) Send(e Email) (interface{}, error) {

	if reflect.DeepEqual(s, SparkPostService{}) {
		return nil, errors.New("SparkPostService struct must be properly initialised")
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

	tx := &sp.Transmission{
		Recipients: []string{e.RecipientEmail},
		Content: sp.Content{
			HTML:    emailHTML,
			From:    fmt.Sprintf("%s <%s>", e.FromName, e.FromEmail),
			Subject: e.Subject,
		},
	}
	if s.apiTesting {
		return tx, nil
	}
	_, response, err := s.client.Send(tx)
	if err != nil {
		return response, err
	}

	return response, nil
}

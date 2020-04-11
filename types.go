package ed

import hermes "github.com/matcornic/hermes/v2"

type Email struct {
	FromName       string
	FromEmail      string
	RecipientEmail string
	Subject        string
	HermesTheme    hermes.Hermes
	HermesEmail    hermes.Email
}

func (e Email) HTML() (string, error) {
	return e.HermesTheme.GenerateHTML(e.HermesEmail)
}

func (e Email) PlainText() (string, error) {
	return e.HermesTheme.GeneratePlainText(e.HermesEmail)
}

package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type MailMessage struct {
	HTMLContent string  `json:"htmlContent"`
	Body        string  `json:"body"`
	Subject     string  `json:"subject"`
	Emails      []Email `json:"emails"`
	MailSender  string  `json:"mailSender"`
	TenantName  string  `json:"labname"`
	EmailExpSys string  `json:"emailExpeditionSystem"`
}

func (m MailMessage) Parse() ([]byte, error) {
	return json.Marshal(m)
}

type Email struct {
	Email string `json:"email"`
}

type EmailNotifier struct {
	url *url.URL
}

func (e *EmailNotifier) Send(message any) error {
	m, ok := message.(MailMessage)
	if !ok {
		return InvalidType
	}

	payload, _ := json.Marshal(m)
	body := bytes.NewBuffer(payload)
	resp, err := http.Post(e.url.Host, "application/json", body)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusUnprocessableEntity {
		return fmt.Errorf("something went wrong")
	}
	return nil
}

func WithEmail(endpoint string) (Notifier, error) {
	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	return &EmailNotifier{
		url: url,
	}, nil
}

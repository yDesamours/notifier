package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SMS struct {
	Id          int    `json:"id"`
	Destination string `json:"destination"`
	Body        string `json:"body"`
	Send_Date   string `json:"send_Date"`
}

type SMSNotifier struct {
	url *url.URL
}

func (s *SMSNotifier) Send(sms any) error {
	m, ok := sms.(SMS)
	if !ok {
		return InvalidType
	}
	payload, _ := json.Marshal(m)
	body := bytes.NewBuffer(payload)
	res, err := http.Post(s.url.Host, "application/json", body)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("someting went wong")
	}
	return nil
}

func WithSMS(endpoint string) (Notifier, error) {
	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &SMSNotifier{
		url: url,
	}, nil

}

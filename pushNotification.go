package notifier

import (
	"log"

	"github.com/appleboy/go-fcm"
)

type FireBaseNotifier struct {
	client *fcm.Client
}

func (f *FireBaseNotifier) Send(m any) error {
	message, ok := m.(FirebaseMessage)
	if !ok {
		return InvalidType
	}

	for _, receiver := range message.TO {
		go func(to string) {
			data := &fcm.Message{Data: message.Data, Notification: message.Notification, To: to}
			_, err := f.client.Send(data)
			if err != nil {
				log.Fatal(err)
			}
		}(receiver)
	}
	return nil
}

func WithFirebase(serverKey string) (Notifier, error) {
	client, err := fcm.NewClient(serverKey)
	if err != nil {
		return nil, err
	}

	return &FireBaseNotifier{client: client}, nil
}

type FirebaseMessage struct {
	TO           []string          `json:"receivers"`
	Data         map[string]any    `json:"data"`
	Notification *fcm.Notification `json:"notification"`
}

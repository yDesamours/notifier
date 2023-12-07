package notify

import "fmt"

type Notifier interface {
	Send(any) error
}

var InvalidType = fmt.Errorf("invalid type")

// type Push interface {
// 	fcm.Message
// }

package useCases

import (
	"errors"
	"fmt"

	"github.com/igloar96/hexa-notification/core/ports"
	outputPort "github.com/igloar96/hexa-notification/core/ports"
)

type CreateNotification struct {
	notificationAdapterList *[]outputPort.OutputPort
}

func NewCreateNotification(outboundAdapters *[]outputPort.OutputPort) *CreateNotification {
	return &CreateNotification{
		notificationAdapterList: outboundAdapters,
	}
}

func (s *CreateNotification) Execute(driver ports.InputPort) []error {
	var errList []error
	message, err := driver.GetMessage()
	if err != nil {
		return append(errList, fmt.Errorf("ERROR WITH NOTIFICATION_DRIVER_ADAPTER DETAILS: %s", err))

	}

	if message.Text == "" {
		return append(errList, errors.New("text is required"))
	}

	for _, adapter := range *s.notificationAdapterList {
		_, err := adapter.Notificate(message)
		if err != nil {
			errList = append(errList, err)
		}
	}

	return errList
}

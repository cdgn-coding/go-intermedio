package main

import "fmt"

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type SMSNotifcation struct {
}

func (sn SMSNotifcation) GetSender() ISender {
	return SMSNotificationSender{}
}

func (sn SMSNotifcation) SendNotification() {
	fmt.Println("Sending notification through SMS")
}

type EmailNotification struct {
}

func (en EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func (en EmailNotification) SendNotification() {
	fmt.Println("Sending notification through Email")
}

type SMSNotificationSender struct {
}

func (S SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (S SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

type EmailNotificationSender struct {
}

func (e EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (e EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func GetNotificationFactory(notificationType string) (INotificationFactory, error) {
	switch notificationType {
	case "SMS":
		return &SMSNotifcation{}, nil
	case "Email":
		return &EmailNotification{}, nil
	default:
		return nil, fmt.Errorf("unhandled notification type")
	}
}

func SendNotification(factory INotificationFactory) {
	factory.SendNotification()
}

func GetMethod(factory INotificationFactory) {
	fmt.Println(factory.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := GetNotificationFactory("SMS")
	emailFactory, _ := GetNotificationFactory("Email")

	SendNotification(smsFactory)
	SendNotification(emailFactory)

	GetMethod(smsFactory)
	GetMethod(emailFactory)
}

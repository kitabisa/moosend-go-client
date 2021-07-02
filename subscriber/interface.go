package subscriber

import "github.com/kitabisa/moosend-go-client/commons"

type Client interface {
	GetSubsriberByEmail(
		format commons.Format,
		mailingListID,
		email string) (returnData Subscriber, err error)
	GetSubscriberByID(
		format commons.Format,
		mailingListID,
		id string) (returnData Subscriber, err error)
	AddSubscriber(
		format commons.Format,
		mailingListID string,
		request SubscribeRequest) (returnData Subscriber, err error)
	UpdateSubscriber(
		format commons.Format,
		mailingListID,
		subscriberID string,
		request SubscribeRequest) (returnData Subscriber, err error)
	UnsubscribeFromAccount(format commons.Format, request UnsubscribeRequest) (err error)
	UnsubscribeFromMailingList(format commons.Format, mailingListID string, request UnsubscribeRequest) (err error)
}

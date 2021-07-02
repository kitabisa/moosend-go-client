package mailing_list

import "github.com/kitabisa/moosend-go-client/commons"

type Client interface {
	GetAllActiveMailingLists(
		format commons.Format,
		withStatistics bool,
		shortBy ShortBy,
		sortMethod SortMethod) (returnData []MailingList, err error)
}

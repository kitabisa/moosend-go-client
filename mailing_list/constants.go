package mailing_list

type ShortBy string

const (
	Name        ShortBy = "Name"
	Subject     ShortBy = "Subject"
	Status      ShortBy = "Status"
	DeliveredOn ShortBy = "DeliveredOn"
	CreatedOn   ShortBy = "CreatedOn"
)

type SortMethod string

const (
	ASC  SortMethod = "ASC"
	DESC SortMethod = "DESC"
)


package models

type Company struct {
	INN   string
	KPP   string
	Owner *Owner
}

type Owner struct {
	Firstname  string
	Middlename string
	Lastname   string
}

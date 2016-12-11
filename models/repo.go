package models

type Repo interface {
	CreateEvent(*Event) error
	SearchEvents(*Query, *[]Event) error
}

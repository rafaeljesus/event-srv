package checker

type api struct{}

func NewApi() *api {
	return &api{}
}

func (a *api) IsAlive() bool {
	return true
}

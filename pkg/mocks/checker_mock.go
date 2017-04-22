package mocks

type CheckerMock struct{}

func NewCheckerMock() *CheckerMock {
	return &CheckerMock{}
}

func (c *CheckerMock) IsAlive() bool {
	return true
}

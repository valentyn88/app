package app

import (
	"errors"
)

type UserServiceMockErr struct{}

func (usm UserServiceMockErr) Get(id int) (*User, error) {
	return nil, errors.New("Error on Get from UserServiceMockErr")
}

func (usm UserServiceMockErr) Update(id int, values map[string]interface{}) error {
	return errors.New("Error on Update from UserServiceMockErr")
}

func (usm UserServiceMockErr) Create(u *User) error {
	return errors.New("Error on Create from UserServiceMockErr")
}

func (usm UserServiceMockErr) Visits() (Visits, error) {
	return nil, errors.New("Error on Visits from UserServiceMockErr")
}

type UseServiceMock struct {
}

func (usm UseServiceMock) Get(id int) (*User, error) {
	return getNewUser(), nil
}

func (usm UseServiceMock) Update(id int, values map[string]interface{}) error {
	return errors.New("Error on Update from UserServiceMockErr")
}

func (usm UseServiceMock) Create(u *User) error {
	return errors.New("Error on Create from UserServiceMockErr")
}

func (usm UseServiceMock) Visits() (Visits, error) {
	return nil, errors.New("Error on Visits from UserServiceMockErr")
}

func getNewUser() *User {
	return &User{ID: 1, Email: "test@gmail.com", FirstName: "FirstNameTest", LastName: "LastNameTest",
		Gender: true, BirthDate: 1988, Age: 29}
}

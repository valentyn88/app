package app

import (
	"errors"
	"fmt"
)

type BasicUser struct {
	cache *Cache
}

func NewBasicUser(cache *Cache) *BasicUser {
	return &BasicUser{
		cache: cache,
	}
}

func (bu BasicUser) Get(id int) (*User, error) {
	if user, ok := bu.cache.Users[id]; ok {
		return user, nil
	}

	return nil, errors.New(fmt.Sprintf("User with ID %d not exists", id))
}

func (bu BasicUser) Update(id int, values map[string]interface{}) error {
	if user, ok := bu.cache.Users[id]; ok {
		user.UpdateFields(values)
		return nil
	}

	return errors.New(fmt.Sprintf("User with ID %d doesn't exists", id))
}

func (bu BasicUser) Create(u *User) error {
	if _, ok := bu.cache.Users[u.ID]; ok {
		return errors.New(fmt.Sprintf("User with ID %d has already exists", u.ID))
	}

	bu.cache.Users[u.ID] = u

	return nil
}

func (bu BasicUser) Visits() (Visits, error) {
	return nil, nil
}

func (u *User) UpdateFields(fields map[string]interface{}) {
	if email, ok := fields["email"]; ok && u.Email != email.(string) {
		u.Email = email.(string)
	}

	if firstName, ok := fields["first_name"]; ok && u.FirstName != firstName.(string) {
		u.FirstName = firstName.(string)
	}

	if lastName, ok := fields["last_name"]; ok && u.LastName != lastName.(string) {
		u.LastName = lastName.(string)
	}

	if gender, ok := fields["gender"]; ok && u.Gender != gender.(bool) {
		u.Gender = gender.(bool)
	}

	if birthDate, ok := fields["birth_date"]; ok && u.BirthDate != int(birthDate.(float64)) {
		u.BirthDate = int(birthDate.(float64))
	}

	if age, ok := fields["age"]; ok && u.Age != int(age.(float64)) {
		u.Age = int(age.(float64))
	}
}

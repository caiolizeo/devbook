package models

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	NickName  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.Format()
	return nil
}

func (user *User) validate() error {
	userValues := reflect.ValueOf(user).Elem()
	userFields := reflect.TypeOf(User{})
	for i := range userValues.NumField() {
		if reflect.DeepEqual(userValues.Field(i).Interface(), reflect.Zero(userValues.Field(i).Type()).Interface()) &&
			(userFields.Field(i).Name != "Id" && userFields.Field(i).Name != "CreatedAt") {
			return fmt.Errorf("o campo %s é obrigatório e não pode ficar em branco", strings.ToLower(userFields.Field(i).Name))
		}
	}

	return nil
}

func (user *User) Format() {
	user.Name = strings.TrimSpace(user.Name)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Email = strings.TrimSpace(user.Email)
}

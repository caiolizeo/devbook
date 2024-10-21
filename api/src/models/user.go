package models

import (
	"api/src/security"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	NickName  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.Format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step string) error {
	userValues := reflect.ValueOf(user).Elem()
	userFields := reflect.TypeOf(*user)

	for i := 0; i < userValues.NumField(); i++ {
		fieldValue := userValues.Field(i)
		fieldType := userFields.Field(i)

		if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) &&
			fieldType.Name != "Id" && fieldType.Name != "CreatedAt" {

			if step != "register" && fieldType.Name == "Password" {
				continue
			}

			return fmt.Errorf("o campo %s é obrigatório e não pode ficar em branco", strings.ToLower(fieldType.Name))
		}
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("o e-mail inserido é inválido")
	}

	return nil
}

func (user *User) Format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPassword)
	}

	return nil
}

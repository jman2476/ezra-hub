package apicaller

import (
	"fmt"
	"reflect"
)

func (c *Client) UpdateUser(nu NewUser) (User, error) {
	v := reflect.ValueOf(nu)
	fields := v.Fields()

	for data, field := range fields {
		if field.Interface() != "" {
			fmt.Printf("\rNew %s: %s\n", data.Name, field.Interface())
		}
	}

	return User{}, nil
}

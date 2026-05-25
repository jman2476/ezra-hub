package apicaller

import "fmt"

func (c *Client) NewUser(signupInfo NewUser) (User, error) {
	user, err := CreateNewResource[User](c, signupInfo, false)
	if err != nil {
		return User{}, err
	}

	fmt.Printf("\rNew user %s created", user.Name)

	var login = UserLogin{
		Name:     user.Name,
		Email:    user.Email,
		Password: signupInfo.Password,
	}

	return c.LoginUser(login)
}

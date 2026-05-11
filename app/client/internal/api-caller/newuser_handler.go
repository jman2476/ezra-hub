package apicaller

func (c *Client) NewUser(signupInfo NewUser) (User, error) {
	user, err := CreateNewResource[User](c, signupInfo)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

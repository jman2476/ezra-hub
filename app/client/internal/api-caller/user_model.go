package apicaller

import "time"

type NewUser struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func (u NewUser) GetLogName() string {
	return "new user"
}

func (u NewUser) GetEndpointURL(c *Client) string {
	return c.baseURL + "/api/users"
}

func (u NewUser) NewEmptyStruct() User {
	return User{}
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type User struct {
	Name          string    `json:"name"`
	PhoneNumber   string    `json:"phone_number"`
	Email         string    `json:"email"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Token         string    `json:"jwt"`
	Refresh       string    `json:"refresh_token"`
	Subscriptions []string  `json:"subs"`
}

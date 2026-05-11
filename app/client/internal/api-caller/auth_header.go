package apicaller

import "net/http"

func (c *Client) MakeAuthHeader() (http.Header, bool) {
	var header = http.Header{}
	if c.token == "" {
		return header, false
	}
	header.Add("Authorization", "Bearer "+c.token)
	return header, true
}

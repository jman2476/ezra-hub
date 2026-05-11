package apicaller

import "net/http"

func (c *Client) MakeAuthHeader(token string) http.Header {
	var header = http.Header{}
	header.Add("Authorization", "Bearer "+token)
	return header
}

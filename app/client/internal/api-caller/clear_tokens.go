package apicaller

import "time"

func (c *Client) ClearTokens() {
	c.token = ""
	c.refresh = ""
	c.lastRefresh = time.Time{}
}

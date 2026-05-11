package apicaller

import (
	"fmt"
)

func (c *Client) NewEvent(newEventInfo NewEvent) (Event, error) {
	resVal, err := c.CreateNewResource(newEventInfo)
	if err != nil {
		return Event{}, err
	}
	newEvent, ok := resVal.(Event)
	if !ok {
		return Event{}, fmt.Errorf("Type Error: response value is not Event struct, it is type %T %T", newEvent, resVal)
	}
	return newEvent, nil
}

package apicaller

func (c *Client) NewEvent(newEventInfo NewEvent) (Event, error) {
	event, err := CreateNewResource[Event](c, newEventInfo)
	if err != nil {
		return Event{}, err
	}

	return event, nil
}

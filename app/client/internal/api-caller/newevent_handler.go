package apicaller

func (c *Client) NewEvent(newEventInfo NewEvent) (EventwName, error) {
	event, err := CreateNewResource[EventwName](c, newEventInfo, false)
	if err != nil {
		return EventwName{}, err
	}

	return event, nil
}

package addyrest

// https://app.addy.io/docs/#failed-deliveries-GETapi-v1-failed-deliveries
func (c *Client) FailedDeliveriesGet() (*FailedDeliveriesWrap, error) {
	return get[FailedDeliveriesWrap](c, "api/v1/failed-deliveries")
}

// https://app.addy.io/docs/#failed-deliveries-GETapi-v1-failed-deliveries--id-
func (c *Client) FailedDeliveryGet(id string) (*FailedDeliveryWrap, error) {
	return get[FailedDeliveryWrap](c, "api/v1/failed-deliveries/"+id)
}

// https://app.addy.io/docs/#failed-deliveries-DELETEapi-v1-failed-deliveries--id-
func (c *Client) FailedDeliveryDel(id string) error {
	_, err := delete[any](c, "api/v1/failed-deliveries/"+id)
	return err
}

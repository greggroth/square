package square

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	AccessToken string
	baseURL     string
}

// Fetch a payment by ID
// https://docs.connect.squareup.com/api/connect/v1/#get-paymentid
func (c *Client) GetPayment(config *GetPaymentRequest) (*Payment, error) {
	var locationId string

	if config.LocationId == "" {
		locationId = "me"
	} else {
		locationId = config.LocationId
	}

	u, err := url.Parse(c.baseURL + locationId + "/payments/" + config.PaymentId)

	if err != nil {
		return nil, err
	}

	payment := &Payment{}
	c.request(u, payment)
	return payment, nil
}

func (c *Client) request(url *url.URL, responseHolder interface{}) error {
	req, err := http.NewRequest("GET", url.String(), nil)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+c.AccessToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(contents, responseHolder)
	return err
}

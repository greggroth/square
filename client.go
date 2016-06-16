package square

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

const iSO8601 = "2006-01-02T15:04:05Z"

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

// Fetch collection of payments.
// https://docs.connect.squareup.com/api/connect/v1/#get-payments
// TODO Respect optional query params
func (c *Client) ListPayments(config *ListPaymentsRequest) (*ListPaymentsResponse, error) {
	var locationId string

	if config.LocationId == "" {
		locationId = "me"
	} else {
		locationId = config.LocationId
	}

	u, err := url.Parse(c.baseURL + locationId + "/payments")

	if err != nil {
		return nil, err
	}

	q := u.Query()

	if !config.BeginTime.IsZero() {
		q.Set("begin_time", config.BeginTime.Format(iSO8601))
	}

	if !config.EndTime.IsZero() {
		q.Set("end_time", config.EndTime.Format(iSO8601))
	}

	if config.Limit != 0 {
		q.Set("limit", strconv.FormatInt(config.Limit, 10))
	}

	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.AccessToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	listPaymentResp := ListPaymentsResponse{}
	err = json.Unmarshal(contents, &listPaymentResp.Payments)

	link := resp.Header.Get("Link")
	re := regexp.MustCompile(`\<([^\>]*)`)
	match := re.FindStringSubmatch(link)

	if len(match) > 0 {
		listPaymentResp.NextPageURL = match[1]
	}

	return &listPaymentResp, err
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

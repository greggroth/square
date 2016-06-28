package square

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPayment(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadFile("testdata/payment.json")

		if err != nil {
			t.Error(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	}))
	defer ts.Close()

	client := Client{AccessToken: "faketoken", baseURL: ts.URL}
	resp, err := client.GetPayment(&GetPaymentRequest{})

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Jq74mCczmFXk1tC10GB", resp.Id)
}

func TestListPayments(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadFile("testdata/payments.json")

		if err != nil {
			t.Error(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Link", "<https://connect.squareup.com/v1/LOCATION_ID/payments?batch_token=BATCH_TOKEN>;rel='next'")
		w.Write(bytes)
	}))
	defer ts.Close()

	client := Client{AccessToken: "faketoken", baseURL: ts.URL}
	resp, err := client.ListPayments(&ListPaymentsRequest{})

	if err != nil {
		t.Error(err)
	}

	assert.Len(t, resp.Payments, 1)
	assert.Equal(t, "Jq74mCczmFXk1tC10GB", resp.Payments[0].Id)
	assert.Equal(t, "https://connect.squareup.com/v1/LOCATION_ID/payments?batch_token=BATCH_TOKEN", resp.NextPageURL)
}

func TestListPaymentsLastPage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadFile("testdata/payments.json")

		if err != nil {
			t.Error(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	}))
	defer ts.Close()

	client := Client{AccessToken: "faketoken", baseURL: ts.URL}
	resp, err := client.ListPayments(&ListPaymentsRequest{})

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "", resp.NextPageURL)
}

func TestListPaymentsWithBeginTime(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadFile("testdata/payments.json")

		if err != nil {
			t.Error(err)
		}

		path, _ := url.QueryUnescape(r.URL.String())
		if path != "/me/payments?begin_time=2016-06-01T08:00:00Z&end_time=2016-06-03T09:30:00Z&limit=50" {
			t.Error("Request was made to:", path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Link", "<https://connect.squareup.com/v1/LOCATION_ID/payments?batch_token=BATCH_TOKEN>;rel='next'")
		w.Write(bytes)
	}))
	defer ts.Close()

	client := Client{AccessToken: "faketoken", baseURL: ts.URL}
	_, err := client.ListPayments(&ListPaymentsRequest{
		BeginTime: time.Date(2016, 6, 1, 8, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2016, 6, 3, 9, 30, 0, 0, time.UTC),
		Limit:     50,
	})

	if err != nil {
		t.Error(err)
	}
}

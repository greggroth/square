package square

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListShifts(t *testing.T) {
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

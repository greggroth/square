## Square Connect v1 API Client

WARNING:  This is very WIP and the API is subject to change.

Go wrapper for the [v1 Square Connect API](https://docs.connect.squareup.com/api/connect/v1/).

### Usage

```go
// Create a new client
client := square.NewClient("yourtoken")

// Fetch a payment by ID
resp, err := client.GetPayment(&GetPaymentRequest{PaymentId: "Jq74mCczmFXk1tC10GB"})
```

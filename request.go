package square

// Request parameters for fetching a payment.  If the LocationId is ommitted
// from the GetPaymentRequest, "me" usedd as the "location_id" in the request.
type GetPaymentRequest struct {
	LocationId string
	PaymentId  string
}

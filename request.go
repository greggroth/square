package square

import "time"

// Request parameters for fetching a payment.  If the LocationId is ommitted
// from the GetPaymentRequest, "me" used as the "location_id" in the request.
type GetPaymentRequest struct {
	LocationId string
	PaymentId  string
}

// Request parameters for fetching a collection of payment.  If the LocationId
// is ommitted from the GetPaymentRequest, "me" used as the "location_id" in
// the request.
type ListPaymentsRequest struct {
	LocationId string
	BeginTime  time.Time
	EndTime    time.Time
	Limit      int64
}

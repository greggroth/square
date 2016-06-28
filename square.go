package square

// Create a new client for working with the Square v1 API.
func NewClient(accessToken string) *Client {
	return &Client{AccessToken: accessToken, baseURL: "https://connect.squareup.com/v1"}
}

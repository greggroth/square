package square

func NewClient(accessToken string) *Client {
	return &Client{AccessToken: accessToken, baseURL: "https://connect.squareup.com/v1/"}
}

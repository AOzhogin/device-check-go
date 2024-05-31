package devicecheck

// Client provides methods to use DeviceCheck API.
type Client struct {
	api  api
	cred Credential
	jwt  jwt
}

// New returns a new DeviceCheck API client.
func New(cred Credential, cfg Config, options ...Option) *Client {

	c := &Client{
		api:  newAPI(cfg.env, options...),
		cred: cred,
		jwt:  newJWT(cfg.issuer, cfg.keyID),
	}

	return c
}

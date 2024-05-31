package devicecheck

import "net/http"

type Option func(a *api)

// WithCustomHttpClient - set your custom http client
func WithCustomHttpClient(httpClient *http.Client) Option {
	return func(a *api) {
		a.client = httpClient
	}
}

// WithCustomBaseURLs - change base url (this option more priority)
func WithCustomBaseURLs(url string) Option {
	return func(a *api) {
		a.baseURL = url
	}
}

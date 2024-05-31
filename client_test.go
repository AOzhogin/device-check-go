package devicecheck

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		cred Credential
		cfg  Config
		want *Client
	}{
		"development": {
			cred: NewCredentialFile("revoked_private_key.p8"),
			cfg:  NewConfig("issuer", "keyID", Development),
			want: &Client{
				api: api{
					client:  http.DefaultClient,
					baseURL: "https://api.development.devicecheck.apple.com/v1",
				},
				cred: credentialFile{
					filename: "revoked_private_key.p8",
				},
				jwt: jwt{
					issuer: "issuer",
					keyID:  "keyID",
				},
			},
		},
		"production": {
			cred: NewCredentialFile("revoked_private_key.p8"),
			cfg:  NewConfig("issuer", "keyID", Production),
			want: &Client{
				api: api{
					client:  http.DefaultClient,
					baseURL: "https://api.devicecheck.apple.com/v1",
				},
				cred: credentialFile{
					filename: "revoked_private_key.p8",
				},
				jwt: jwt{
					issuer: "issuer",
					keyID:  "keyID",
				},
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := New(c.cred, c.cfg)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want '%+v', got '%+v'", c.want, got)
			}
		})
	}
}

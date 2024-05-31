package devicecheck

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestWithCustomBaseURLs(t *testing.T) {

	creds := NewCredentialString("test creds")
	cfg := NewConfig("test issuer", "test keyid", Development)

	type args struct {
		Options []Option
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Changed url", args: args{Options: []Option{WithCustomBaseURLs("https://google.com/api/v1/")}}, want: "https://google.com/api/v1/"},
		{name: "Changed url", args: args{Options: nil}, want: developmentBaseURL},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			client := New(creds, cfg, tt.args.Options...)

			if got := client.api.baseURL; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCustomBaseURLs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithCustomHttpClient(t *testing.T) {

	httpClient := newMockHTTPClient(&http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("success")),
	})

	creds := NewCredentialString("test creds")
	cfg := NewConfig("test issuer", "test keyid", Development)

	type args struct {
		Options []Option
	}
	tests := []struct {
		name string
		args args
		want *http.Client
	}{
		{name: "changed http client", args: args{Options: []Option{WithCustomHttpClient(httpClient)}}, want: httpClient},
		{name: "default http client", args: args{Options: nil}, want: http.DefaultClient},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := New(creds, cfg, tt.args.Options...)

			if got := client.api.client; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCustomHttpClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

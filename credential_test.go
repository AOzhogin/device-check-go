package devicecheck

import (
	"os"
	"reflect"
	"testing"
)

func TestNewCredentialFile(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		filename string
		want     credentialFile
	}{
		"valid filename": {
			filename: "./tests/revoked_private_key.p8",
			want: credentialFile{
				filename: "./tests/revoked_private_key.p8",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := NewCredentialFile(c.filename)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want '%+v', got '%+v'", c.want, got)
			}
		})
	}
}

func TestCredentialFile_key(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		cred  credentialFile
		noErr bool
	}{
		"valid credential": {
			cred: credentialFile{
				filename: "./tests/revoked_private_key.p8",
			},
			noErr: true,
		},
		"invalid credential": {
			cred: credentialFile{
				filename: "credential_test.go",
			},
			noErr: false,
		},
		"unknown filename": {
			cred: credentialFile{
				filename: "./tests/unknown_file.p8",
			},
			noErr: false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			key, err := c.cred.key()

			if c.noErr {
				if err != nil {
					t.Errorf("want 'nil', got '%+v'", err)
				}
				if key == nil {
					t.Error("want 'not nil', got 'nil'")
				}
			} else {
				if err == nil {
					t.Error("want 'not nil', got 'nil'")
				}
				if key != nil {
					t.Errorf("want 'nil', got '%+v'", key)
				}
			}
		})
	}
}

func TestNewCredentialBytes(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		filename string
	}{
		"valid filename": {
			filename: "./tests/revoked_private_key.p8",
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			raw, err := os.ReadFile(c.filename)
			if err != nil {
				t.Errorf("want 'nil', got '%+v'", err)
			}

			got := NewCredentialBytes(raw)
			want := credentialBytes{raw: raw}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("want '%+v', got '%+v'", want, got)
			}
		})
	}
}

func TestCredentialBytes_key(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		filename string
		noErr    bool
	}{
		"valid filename": {
			filename: "./tests/revoked_private_key.p8",
			noErr:    true,
		},
		"invalid private key": {
			filename: "./tests/invalid_private_key.p8",
			noErr:    false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			raw, err := os.ReadFile(c.filename)
			if err != nil {
				t.Errorf("want 'nil', got '%+v'", err)
			}

			cred := NewCredentialBytes(raw)
			key, err := cred.key()

			if c.noErr {
				if err != nil {
					t.Errorf("want 'nil', got '%+v'", err)
				}
				if key == nil {
					t.Error("want 'not nil', got 'nil'")
				}
			} else {
				if err == nil {
					t.Error("want 'not nil', got 'nil'")
				}
				if key != nil {
					t.Errorf("want 'nil', got '%+v'", key)
				}
			}
		})
	}
}

func TestNewCredentialString(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		filename string
	}{
		"valid filename": {
			filename: "./tests/revoked_private_key.p8",
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			raw, err := os.ReadFile(c.filename)
			if err != nil {
				t.Errorf("want 'nil', got '%+v'", err)
			}

			got := NewCredentialString(string(raw))
			want := credentialString{str: string(raw)}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("want '%+v', got '%+v'", want, got)
			}
		})
	}
}

func TestCredentialString_key(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		filename string
		noErr    bool
	}{
		"valid credential": {
			filename: "./tests/revoked_private_key.p8",
			noErr:    true,
		},
		"invalid credential": {
			filename: "credential_test.go",
			noErr:    false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			raw, err := os.ReadFile(c.filename)
			if err != nil {
				t.Errorf("want 'nil', got '%+v'", err)
			}

			cred := NewCredentialString(string(raw))
			key, err := cred.key()

			if c.noErr {
				if err != nil {
					t.Errorf("want 'nil', got '%+v'", err)
				}
				if key == nil {
					t.Error("want 'not nil', got 'nil'")
				}
			} else {
				if err == nil {
					t.Error("want 'not nil', got 'nil'")
				}
				if key != nil {
					t.Errorf("want 'nil', got '%+v'", key)
				}
			}
		})
	}
}

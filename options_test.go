package gocardless

import "testing"

func TestEndpointForToken(t *testing.T) {
	tests := []struct {
		token   string
		wantErr bool
	}{
		{"", true},
		{"dummy_token", false},
	}

	for _, tt := range tests {
		_, err := NewConfig(tt.token)
		if tt.wantErr {
			if err == nil {
				t.Fatalf("expected error if token is invalid, got nil for token: %v", tt.token)
			}
			continue
		}

		if err != nil {
			t.Fatalf("unexepcted error if token is valid, got token: %v, err: %v", tt.token, err)
		}
	}
}

func TestEndpointForConfig(t *testing.T) {
	tests := []struct {
		endpoint string
		wantErr  bool
	}{
		{"http://foo.com", false},
		{"1http://foo.com", true},
	}

	for _, tt := range tests {
		token := "dummy_token"
		_, err := NewConfig(token, WithEndpoint(tt.endpoint))
		if tt.wantErr {
			if err == nil {
				t.Fatalf("expected error if endpoint is invalid, got nil for endpoint: %v", tt.endpoint)
			}
			continue
		}

		if err != nil {
			t.Fatalf("unexepcted error if endpoint is valid, got endpoint: %v, err: %v", tt.endpoint, err)
		}
	}
}

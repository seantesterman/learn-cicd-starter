package auth

import (
	"net/http"
	"testing"
)

func TestAPIKey(t *testing.T) {

	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{name: "empty header", want: "", wantErr: true},
		{name: "valid authorization", want: "123456", wantErr: false},
		{name: "invalid format", want: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := http.Header{}

			if tt.name == "valid authorization" {
				header.Add("Authorization", "ApiKey 123456")
			}

			if tt.name == "invalid format" {
				header.Add("Authorization", "BadFormat") // no "Bearer" prefix
			}

			t.Logf("Debug - Header value %q", header.Get("Authorization"))
			got, err := GetAPIKey(header)
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

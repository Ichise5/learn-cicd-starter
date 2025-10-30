package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		headers http.Header
		want    string
		wantErr bool
	}{
		{name: "noAPIGiven", headers: http.Header{"Authorization": {"some_text"}}, want: "", wantErr: true},
		{name: "noAuthorization", headers: http.Header{"Content-Type": {"application/json"}}, want: "", wantErr: true},
		{name: "correctAuthorization", headers: http.Header{"Authorization": {"ApiKey 1s239g4wvfhdtefertn"}}, want: "1s239g4wvfhdtefertn", wantErr: false},
		{name: "alsoCorrectAuthorization", headers: http.Header{"Authorization": {"ApiKey thisIsValidApiKey"}}, want: "thisIsValidApiKey", wantErr: false},
		{name: "badCase", headers: http.Header{"Authorization": {"APIKEY thisIsAlsoApiKey"}}, want: "", wantErr: true},
		{name: "malformedAPI", headers: http.Header{"Authorization": {"thisIsnotApiKey  right"}}, want: "", wantErr: true},
		{name: "emptyApi", headers: http.Header{"Authorization": {"ApiKey"}}, want: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := auth.GetAPIKey(tt.headers)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetAPIKey() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetAPIKey() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

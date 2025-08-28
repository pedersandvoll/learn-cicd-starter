package auth_test

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "No Authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
		{
			name: "Malformed Authorization header - missing ApiKey prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer something"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Malformed Authorization header - ony ApiKey prefix",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Valid Authorization header",
			headers: http.Header{
				"Authorization": []string{"ApiKey 12345"},
			},
			want:    "12345",
			wantErr: false,
		},
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
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

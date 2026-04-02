package auth

import (
    "errors"
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    tests := []struct {
        name    string
        headers http.Header
        wantKey string
        wantErr error
    }{
        {
            name:    "valid ApiKey header",
            headers: http.Header{"Authorization": {"ApiKey 12345"}},
            wantKey: "12345",
            wantErr: nil,
        },
        {
            name:    "missing Authorization header",
            headers: http.Header{},
            wantKey: "",
            wantErr: ErrNoAuthHeaderIncluded,
        },
        {
            name:    "malformed header wrong prefix",
            headers: http.Header{"Authorization": {"Bearer 12345"}},
            wantKey: "",
            wantErr: errors.New("malformed authorization header"),
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := GetAPIKey(tt.headers)

            if got != tt.wantKey {
                t.Errorf("GetAPIKey() got key %q, want %q", got, tt.wantKey)
            }

            if tt.wantErr == nil {
                if err != nil {
                    t.Errorf("GetAPIKey() unexpected error = %v", err)
                }
            } else {
                if err == nil || err.Error() != tt.wantErr.Error() {
                    t.Errorf("GetAPIKey() error = %v, want %v", err, tt.wantErr)
                }
            }
        })
    }
}


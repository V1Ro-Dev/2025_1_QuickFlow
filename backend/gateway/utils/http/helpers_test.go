package http_test

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"quickflow/gateway/internal/delivery/http/forms"
	errors2 "quickflow/gateway/internal/errors"
	http2 "quickflow/gateway/utils/http"
)

func TestWriteJSONError_Table(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantCode int
		wantErr  string
		wantMsg  string
	}{
		{
			name:     "bad request",
			err:      errors2.New("BAD_REQUEST", "bad", http.StatusBadRequest),
			wantCode: http.StatusBadRequest,
			wantErr:  "BAD_REQUEST",
			wantMsg:  "bad",
		},
		{
			name:     "not found",
			err:      errors2.New("NOT_FOUND", "not found", http.StatusNotFound),
			wantCode: http.StatusNotFound,
			wantErr:  "NOT_FOUND",
			wantMsg:  "not found",
		},
		{
			name:     "internal error",
			err:      errors2.New("INTERNAL", "fail", http.StatusInternalServerError),
			wantCode: http.StatusInternalServerError,
			wantErr:  "INTERNAL",
			wantMsg:  "fail",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			http2.WriteJSONError(rec, tt.err)

			res := rec.Result()
			defer res.Body.Close()

			require.Equal(t, tt.wantCode, res.StatusCode)
			require.Equal(t, "application/json", res.Header.Get("Content-Type"))

			var resp forms.ErrorForm
			require.NoError(t, json.NewDecoder(res.Body).Decode(&resp))
			require.Equal(t, tt.wantErr, resp.ErrorCode)
			require.Equal(t, tt.wantMsg, resp.Message)
		})
	}
}

func createMultipartRequest(t *testing.T, fieldName string, files map[string]string) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for filename, content := range files {
		part, err := writer.CreateFormFile(fieldName, filename)
		require.NoError(t, err)
		_, err = part.Write([]byte(content))
		require.NoError(t, err)
	}

	require.NoError(t, writer.Close())

	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	require.NoError(t, req.ParseMultipartForm(15<<20))

	return req
}

func TestGetFiles_Table(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		files       map[string]string
		wantNames   []string
		expectError bool
	}{
		{
			"two valid files",
			"docs",
			map[string]string{"a.txt": "a", "b.md": "b"},
			[]string{"a.txt", "b.md"},
			false,
		},
		{
			"no files in field",
			"none",
			map[string]string{},
			nil,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := createMultipartRequest(t, tt.field, tt.files)
			files, err := http2.GetFiles(req, tt.field)

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if len(tt.wantNames) == 0 {
					require.Len(t, files, 0)
				} else {
					var actualNames []string
					for _, f := range files {
						actualNames = append(actualNames, f.Name)
					}
					require.ElementsMatch(t, tt.wantNames, actualNames)
				}
			}
		})
	}
}

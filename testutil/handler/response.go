package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func AssertResponse(t *testing.T, got *http.Response, wantStatusCode int, wantJSONBody []byte) {
	t.Helper()

	t.Cleanup(func() { _ = got.Body.Close() })
	gotJSONBody, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatalf("failed to read got body %q: %v", got.Body, err)
	}

	if got.StatusCode != wantStatusCode {
		t.Fatalf("want status code %d, but got %d, body: %q",
			wantStatusCode, got.StatusCode, gotJSONBody)
	}

	// レスポンスボディが必要ない場合
	if len(wantJSONBody) == 0 && len(gotJSONBody) == 0 {
		return
	}

	assertJSON(t, wantJSONBody, gotJSONBody)
}

func assertJSON(t *testing.T, wantJSON, gotJSON []byte) {
	t.Helper()

	var want, got any
	if err := json.Unmarshal(gotJSON, &got); err != nil {
		t.Fatalf("failed to unmarshal got %q: %v", got, err)
	}
	if err := json.Unmarshal(wantJSON, &want); err != nil {
		t.Fatalf("failed to unmarshal want %q: %v", want, err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("got differs: (-got +want)\n%s", diff)
	}
}

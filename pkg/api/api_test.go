package api

import (
	"censor/pkg/censor"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestContext struct {
	api *API
}

func setup(_ *testing.T) TestContext {
	api := New(censor.New())
	return TestContext{api}
}

func TestCommentsValidateWithGoodContent(t *testing.T) {
	ctx := setup(t)

	req_json := `
	{"content":"Test"}
	`
	req := httptest.NewRequest(http.MethodPost, "/comments/validate", strings.NewReader(req_json))
	resp := httptest.NewRecorder()
	ctx.api.router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatal("Wrong status code")
	}
}

func TestCommentsValidateWithBadContent(t *testing.T) {
	ctx := setup(t)

	req_json := `
	{"content":"qwerty"}
	`
	req := httptest.NewRequest(http.MethodPost, "/comments/validate", strings.NewReader(req_json))
	resp := httptest.NewRecorder()
	ctx.api.router.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatal("Wrong status code")
	}
}

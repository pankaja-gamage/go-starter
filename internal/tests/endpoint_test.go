package tests

import (
	"fmt"
	"go-starter/internal/common/data"
	"go-starter/internal/routers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	test *Test
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	test = New()

	defer data.CleanUp()

	return m.Run()
}

func TestHealthCheck(t *testing.T) {
	ts := httptest.NewServer(routers.SetupRouter())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/health", ts.URL))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("expected status code 200, got %v", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := `{"status":"OK"}`
	if string(respBody) != expected {
		t.Errorf("rest endpoint returned unexpected body: got %v want %v",
			string(respBody), expected)
	}
}

func TestAddUserSuccess(t *testing.T) {
	ts := httptest.NewServer(routers.SetupRouter())

	defer ts.Close()

	var request = `{"name": "John Doe","email":"johndoe@gmail.com","mobile":"+1232342233"}`

	resp, err := http.Post(fmt.Sprintf("%s/users", ts.URL), "application/json", strings.NewReader(request))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.StatusCode != 201 {
		t.Fatalf("expected status code 201, got %v", resp.StatusCode)
	}
}

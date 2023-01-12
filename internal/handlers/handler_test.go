package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedstatuscode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := GetRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)

			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedstatuscode {
				t.Errorf("for %s ,expected %d but got %d", e.name, e.expectedstatuscode, resp.StatusCode)
			}
		} else {

		}
	}

}

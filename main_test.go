package main

import (
	"github.com/infoblox/parsecsv"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkProcessCSV(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parsecsv.ProcessCSV()
	}
}

func TestSearch(t *testing.T) {
	r := request(t, "/George") // the endpoint will be the name we want to search inside the address book
	// It will return the whole json string including all information stored for that record.
	rr := httptest.NewRecorder()
	parsecsv.Search(rr, r)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if rr.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}

}

func BenchmarkSearch(b *testing.B) {
	r := request(b, "/George")
	for i := 0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		parsecsv.Search(rw, r)
	}

}

func request(t testing.TB, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

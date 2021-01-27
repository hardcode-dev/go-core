package api

import (
	"encoding/json"
	"go-core/22-extra-movies/pkg/models"
	"go-core/22-extra-movies/pkg/storage/mem"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var api *API

func TestMain(m *testing.M) {
	db := new(mem.DB)
	api = New(db)
	os.Exit(m.Run())
}

func TestAPI_movies(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/movies", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	t.Log("Response: ", rr.Body)
	b, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	var data []models.Movie
	err = json.Unmarshal(b, &data)
	if err != nil {
		t.Fatal(err)
	}
	got := len(data)
	want := 2
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}

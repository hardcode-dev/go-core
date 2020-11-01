package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gosearch/pkg/engine"
	"gosearch/pkg/index/hash"
	"gosearch/pkg/storage/memstore"

	"github.com/gorilla/mux"
)

var api *Service

func TestMain(m *testing.M) {
	router := mux.NewRouter()
	index := hash.New()
	storage := memstore.New()
	engine := engine.New(index, storage)
	api = New(router, engine)
	api.endpoints()

	os.Exit(m.Run())
}

func TestService_Search(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/search/go", nil)

	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	t.Logf("Длина ответа от сервера %d байт", rr.Body.Len())
}

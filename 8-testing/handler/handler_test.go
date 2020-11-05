package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var api API

func TestAPI_ProductsHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/HANDLER", nil)

	rr := httptest.NewRecorder()
	api.ProductsHandler(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	resp := rr.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var data []product
	json.Unmarshal(body, &data)
	t.Logf("Ответ сервера:\n%v\n", data)
}

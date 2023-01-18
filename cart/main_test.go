package main

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

var (
	// mockDB = map[string]*produk{
	// "0":{kodeProduk:"121","namaProduk":"namaProduk","kuantitas":"23"},
	// }
	userJSON = `[{"kodeProduk":"121","namaProduk":"namaProduk","kuantitas":"23"}]`
)

// func executeRequest(req *http.Request) *httptest.ResponseRecorder {
//     rr := httptest.NewRecorder()
//     a.Router.ServeHTTP(rr, req)

//     return rr
// // }
// func produkIndex() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		data := []*produk{
// 			{

// 				KodeProduk: 121,
// 				Nama:       "namaProduk",
// 				Kuantitas:  23,
// 			},
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		if err := json.NewEncoder(w).Encode(data); err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte("Error reading bookmarks"))
// 		}
// 	})
// }
func TestAddData(t *testing.T) {
	data := url.Values{}
	data.Set("kodeProduk", "121")
	data.Set("namaProduk", "namaProdu")
	data.Set("kuantitas", "23")
	encodedata := data.Encode()
	res, _ := http.Post("http://127.0.0.1:5000/api/add", "application/x-www-form-urlencoded", strings.NewReader(encodedata))
	body, err := io.ReadAll(res.Body)
	t.Log("http://127.0.0.1:5000/api/add")
	t.Log(encodedata)
	t.Log(string(body))
	if err != nil {
		t.Errorf("Expected nil, received %s ", err.Error())
	}
	if res.StatusCode != http.StatusCreated {
		// e := echo.New()
		// req := httptest.NewRequest(http.MethodGet, "/", nil)
		// rec := httptest.NewRecorder()
		// c := e.NewContext(req, rec)
		// c.SetPath("/api/get")
		// c.SetParamNames("email")
		// c.SetParamValues("jon@labstack.com")
		// h := &produk{
		// 	KodeProduk: 100,
		// 	Nama:       "Baju hitam",
		// 	Kuantitas:  20,
		// 	Message:    "",
		// }

		// Assertions
		// if assert.NoError(t, h.gatdata(c)) {
		// 	assert.Equal(t, http.StatusOK, rec.Code)
		// 	// assert.Equal(t, userJSON, rec.Body.String())
		// }
		t.Errorf("Expected %d, received %d", http.StatusCreated, res.StatusCode)
	}

}

func TestGetCart(t *testing.T) {
	// Setup
	res, err := http.Get("http://127.0.0.1:5000/api/get")
	body, err := io.ReadAll(res.Body)
	t.Log("http://127.0.0.1:5000/api/get")
	t.Log(string(body))
	if err != nil {
		t.Errorf("Expected nil, received %s ", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}

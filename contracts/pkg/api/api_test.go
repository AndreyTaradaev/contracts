// маршрутизатор для http сервера API приложения Gateway

package api

import (
	"bytes"
	_ "gateway/censor/docs"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI_endpoints(t *testing.T) {
	api, err := New()
	if err != nil {
		t.Fatal(err)
		return
	}
	var testBody1 = []byte(`  {
	   		"content": "Ghbrjkmyj1112 asdasdqertyasdQWEYasd",
	   		"idnews" : 65222,
	   		"authorid":  1121 ,
	   		"time":  5556522222,

	   	"Parent": 46


	   	}`)
	var testBody2 = []byte(` {		
		"content": "Ghbrjkmyj1112 asdasdqertyasdQWErtYasd",
		"idnews" : 65222,		
		"authorid":  1121 ,
		"time":  5556522222,		
		"Parent": 46			
	}`)

	req := httptest.NewRequest(http.MethodPost, "/comment", bytes.NewBuffer(testBody1))
	rr := httptest.NewRecorder()
	api.Router().ServeHTTP(rr, req)
	// Проверяем код ответа.
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	req = httptest.NewRequest(http.MethodPost, "/comment", bytes.NewBuffer(testBody2))
	rr = httptest.NewRecorder()
	api.Router().ServeHTTP(rr, req)
	// Проверяем код ответа.
	if !(rr.Code == http.StatusForbidden) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusForbidden)
	}

	req = httptest.NewRequest(http.MethodPost, "/word/Запрещено", nil)
	rr = httptest.NewRecorder()
	api.Router().ServeHTTP(rr, req)
	// Проверяем код ответа.
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

}

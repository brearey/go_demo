package inet

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCurl(t *testing.T) {
	// Создаем тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "test response"}`))
	}))
	defer server.Close()

	// Тестируем успешный запрос
	body, err := Curl(server.URL)
	if err != nil {
		t.Errorf("Curl вернул ошибку: %v", err)
	}

	expected := `{"message": "test response"}`
	if body != expected {
		t.Errorf("ожидалось %q, получено %q", expected, body)
	}
}

func TestCurl_Error(t *testing.T) {
	// Сервер возвращает 404
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	_, err := Curl(server.URL)
	if err == nil {
		t.Error("ожидалась ошибка для статуса 404")
	}
}

func TestCurl_InvalidURL(t *testing.T) {
	_, err := Curl("invalid-url")
	if err == nil {
		t.Error("ожидалась ошибка для невалидного URL")
	}
}
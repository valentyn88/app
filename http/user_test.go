package http

import (
	"github.com/gorilla/mux"
	"github.com/valentyn88/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_GetUser_WrongID(t *testing.T) {
	h := NewUserHandler(app.UserServiceMockErr{})
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/user/test", nil)
	h.GetUser(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status code: got : %v want %v", w.Code, http.StatusBadRequest)
	}
}

func TestUserHandler_GetUser_NotExistsID(t *testing.T) {
	h := NewUserHandler(app.UserServiceMockErr{})
	req, err := http.NewRequest("GET", "/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/user/{id:[0-9]+}", h.GetUser).Methods("GET")
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("wrong status code: got : %v want %v", rr.Code, http.StatusBadRequest)
	}
}

func TestNewUserHandler_GetUser_Success(t *testing.T) {
	h := NewUserHandler(app.UseServiceMock{})
	req, err := http.NewRequest("GET", "/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/user/{id:[0-9]+}", h.GetUser).Methods("GET")
	router.ServeHTTP(rr, req)

	expected := `{"id":1,"email":"test@gmail.com","first_name":"FirstNameTest","last_name":"LastNameTest","gender":true,"birth_date":1988,"age":29}`
	if rr.Code != http.StatusOK && rr.Body.String() != expected {
		t.Errorf("handler return unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}

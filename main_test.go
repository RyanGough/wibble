package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakeDB struct{}

func (f FakeDB) GetName(id int) (string, error) {
	if id != 1234 {
		return "", errors.New("blah")
	}
	return "Ray Purchase", nil
}

func TestInvalidRequestReturns400(t *testing.T) {
	request, _ := http.NewRequest("GET", "/name?id=wibble", nil)
	response := httptest.NewRecorder()
	e := Engine(FakeDB{})
	e.ServeHTTP(response, request)
	if response.Code != 400 {
		t.Errorf("Expected 400, got %v", response.Code)
	}
	body, _ := ioutil.ReadAll(response.Body)
	if string(body) != "id should be int" {
		t.Errorf("Unexpected error message, got %v", string(body))
	}
}

func TestValidRequestReturnsName(t *testing.T) {
	request, _ := http.NewRequest("GET", "/name?id=1234", nil)
	response := httptest.NewRecorder()
	e := Engine(FakeDB{})
	e.ServeHTTP(response, request)
	if response.Code != 200 {
		t.Errorf("Expected 200, got %v", response.Code)
	}
	body, _ := ioutil.ReadAll(response.Body)
	if string(body) != "Ray Purchase" {
		t.Errorf("Unexpected name, got %v", string(body))
	}
}

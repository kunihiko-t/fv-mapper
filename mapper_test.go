package fvm

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestHandler struct {
}

type TestSequentialHandler struct {
}

type TestCamelHandler struct {
}

type TestSnakeHandler struct {
}

var errs []error
var errsSeq []error
var errsCamel []error
var errsSnake []error

func (h *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	m := GetMap(r)

	if m["name_1"] != "test_1" {
		e := fmt.Sprintf("Caught %v expected : test_1", m["name_1"])
		errs = append(errs, errors.New(e))
	}

	if m["name_2"] != "test_2" {
		e := fmt.Sprintf("Caught %v expected : test_2", m["name_2"])
		errs = append(errs, errors.New(e))
	}

	if m["name_3"] != "test_3" {
		e := fmt.Sprintf("Caught %v expected : test_3", m["name_3"])
		errs = append(errs, errors.New(e))
	}

	if m["another_name_1"] != "a" {
		e := fmt.Sprintf("Caught %v expected : a", m["another_name_1"])
		errs = append(errs, errors.New(e))
	}

	//TODO More test case like a test[]
}

func (h *TestSequentialHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	m := GetMapSequential("name", r)

	if m["name_1"] != "test_1" {
		e := fmt.Sprintf("Caught %v expected : test_1", m["name_1"])
		errsSeq = append(errsSeq, errors.New(e))
	}

	if m["name_2"] != "test_2" {
		e := fmt.Sprintf("Caught %v expected : test_2", m["name_2"])
		errsSeq = append(errsSeq, errors.New(e))
	}

	if m["name_3"] != "test_3" {
		e := fmt.Sprintf("Caught %v expected : test_3", m["name_3"])
		errsSeq = append(errsSeq, errors.New(e))
	}

	if m["another_name_1"] != "" {
		e := fmt.Sprintf("Caught %v expected : blank", m["another_name_1"])
		errsSeq = append(errsSeq, errors.New(e))
	}

}

func (h *TestCamelHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	m := GetCamelMap(false, r)

	if m["nameTest"] != "foo" {
		e := fmt.Sprintf("Caught %v expected : foo", m["nameTest"])
		errsCamel = append(errsCamel, errors.New(e))
	}

	if m["name_test"] != "" {
		e := fmt.Sprintf("Caught %v expected : blank", m["name_test"])
		errsCamel = append(errsCamel, errors.New(e))
	}

	if m["nameA"] != "bar" {
		e := fmt.Sprintf("Caught %v expected : bar", m["nameA"])
		errsCamel = append(errsCamel, errors.New(e))
	}

	if m["name_a"] != "" {
		e := fmt.Sprintf("Caught %v expected : blank", m["name_a"])
		errsCamel = append(errsCamel, errors.New(e))
	}

	m = GetCamelMap(true, r)

	if m["NameTest"] != "foo" {
		e := fmt.Sprintf("Caught %v expected : foo", m["nameTest"])
		errsCamel = append(errsCamel, errors.New(e))
	}

	if m["name_test"] != "" {
		e := fmt.Sprintf("Caught %v expected : blank", m["name_test"])
		errsCamel = append(errsCamel, errors.New(e))
	}

	if m["NameA"] != "bar" {
		e := fmt.Sprintf("Caught %v expected : bar", m["nameA"])
		errsCamel = append(errsCamel, errors.New(e))
	}

	if m["name_a"] != "" {
		e := fmt.Sprintf("Caught %v expected : blank", m["name_a"])
		errsCamel = append(errsCamel, errors.New(e))
	}

	if m["あいう"] != "え" {
		e := fmt.Sprintf("Caught %v expected : え", m["あいう"])
		errsCamel = append(errsCamel, errors.New(e))
	}

}

func (h *TestSnakeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	m := GetSnakeMap(r)

	if m["name_test"] != "foo" {
		e := fmt.Sprintf("Caught %v expected : foo", m["name_test"])
		errsSnake = append(errsSnake, errors.New(e))
	}

	if m["nameTest"] != "" {
		e := fmt.Sprintf("Caught %v expected : blank", m["nameTest"])
		errsSnake = append(errsSnake, errors.New(e))
	}

	if m["name_a"] != "bar" {
		e := fmt.Sprintf("Caught %v expected : bar", m["name_a"])
		errsSnake = append(errsSnake, errors.New(e))
	}

	if m["nameA"] != "" {
		e := fmt.Sprintf("Caught %v expected : blank", m["nameA"])
		errsSnake = append(errsSnake, errors.New(e))
	}

}

func TestGetMap(t *testing.T) {

	errs = []error{}
	h := &TestHandler{}
	ts := httptest.NewServer(h)
	defer ts.Close()
	request := gorequest.New()

	//POST
	_, _, err := request.Post(ts.URL).Type("form").
		Send(`{ "name_1": "test_1","name_2": "test_2", "name_3": "test_3", "another_name_1": "a" }`).
		End()

	if err != nil {
		t.Error("unexpected error:", err)
	}

	for _, e := range errs {
		t.Errorf("[POST ]%v", e)
	}

	//GET
	errs = []error{}
	_, _, err = request.Get(ts.URL).
		Query(`{ "name_1": "test_1","name_2": "test_2", "name_3": "test_3", "another_name_1": "a" }`).
		End()
	if err != nil {
		t.Error("unexpected error:", err)
	}

	for _, e := range errs {
		t.Errorf("[GET] %v", e)
	}

}

func TestGetMapSequential(t *testing.T) {

	errsSeq = []error{}
	h := &TestSequentialHandler{}
	ts := httptest.NewServer(h)
	defer ts.Close()
	request := gorequest.New()

	//POST
	_, _, err := request.Post(ts.URL).Type("form").
		Send(`{ "name_1": "test_1","name_2": "test_2", "name_3": "test_3", "another_name_1": "a" }`).
		End()

	if err != nil {
		t.Error("unexpected error:", err)
	}

	for _, e := range errsSeq {
		t.Errorf("[POST ]%v", e)
	}

	//GET
	errsSeq = []error{}
	_, _, err = request.Get(ts.URL).
		Query(`{ "name_1": "test_1","name_2": "test_2", "name_3": "test_3", "another_name_1": "a" }`).
		End()
	if err != nil {
		t.Error("unexpected error:", err)
	}

	for _, e := range errsSeq {
		t.Errorf("[GET] %v", e)
	}

}

func TestGetCamelMap(t *testing.T) {

	errsCamel = []error{}
	h := &TestCamelHandler{}
	ts := httptest.NewServer(h)
	defer ts.Close()
	request := gorequest.New()

	//POST
	_, _, err := request.Post(ts.URL).Type("form").
		Send(`{ "name_test": "foo","nameA": "bar","あ_い_う": "え" }`).
		End()

	if err != nil {
		t.Error("unexpected error:", err)
	}

	for _, e := range errsCamel {
		t.Errorf("[POST ]%v", e)
	}

	//GET
	errsCamel = []error{}
	_, _, err = request.Get(ts.URL).
		Query(`{ "name_test": "foo","nameA": "bar","あ_い_う": "え" }`).
		End()
	if err != nil {
		t.Error("unexpected error:", err)
	}

	for _, e := range errsCamel {
		t.Errorf("[GET] %v", e)
	}

}

func TestGetSnakeMap(t *testing.T) {

	errsSnake = []error{}
	h := &TestSnakeHandler{}
	ts := httptest.NewServer(h)
	defer ts.Close()
	request := gorequest.New()

	//POST
	_, _, err := request.Post(ts.URL).Type("form").
		Send(`{ "name_test": "foo","nameA": "bar" }`).
		End()

	if err != nil {
		t.Error("unexpected error:", err)
	}

	for _, e := range errsSnake {
		t.Errorf("[POST ]%v", e)
	}

	//GET
	errsSnake = []error{}
	_, _, err = request.Get(ts.URL).
		Query(`{ "name_test": "foo","nameA": "bar" }`).
		End()
	if err != nil {
		t.Error("unexpected error:", err)
	}

	for _, e := range errsSnake {
		t.Errorf("[GET] %v", e)
	}

}

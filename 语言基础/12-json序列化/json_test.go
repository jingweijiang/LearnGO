package main

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	"testing"
)

type Student struct {
	Name string
	Age  int

	Gender bool
}

type Class struct {
	ID       string
	Students []Student
}

var (
	s = Student{
		Name:   "Tom",
		Age:    20,
		Gender: true,
	}
	c = Class{
		ID:       "1001",
		Students: []Student{s, s, s, s}}
)

// go test .\json_test.go -v
// go test .\json_test.go -v -run TestJson
// go test .\json_test.go -v -run TestSonicJson
func TestJson(t *testing.T) {
	jsonStr, err := json.Marshal(c)
	if err != nil {
		t.Fail()
	}

	var c2 Class
	err = json.Unmarshal(jsonStr, &c2)
	if err != nil {
		t.Fail()
	}

	if c2.ID != c.ID || len(c2.Students) != len(c.Students) {
		t.Fail()
	}

}

func TestSonicJson(t *testing.T) {
	jsonStr, err := sonic.Marshal(c)
	if err != nil {
		t.Fail()
	}

	var c2 Class
	err = sonic.Unmarshal(jsonStr, &c2)
	if err != nil {
		t.Fail()
	}

	if c2.ID != c.ID || len(c2.Students) != len(c.Students) {
		t.Fail()
	}
}

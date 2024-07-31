package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestGetStudentInfo(t *testing.T) {
	stuid := "1001"
	stu := GetStudentInfo(stuid)

	if len(stu.Name) == 0 {
		t.Fail()
	} else {
		fmt.Printf("%+v\n", stu)
	}

}

func TestGetName(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:2345/get_name?studentId=1001")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		} else {
			fmt.Println(string(body))
		}
	}
}

func TestGetAge(t *testing.T) {
	resp, err := http.PostForm("http://127.0.0.1:2345/get_age", url.Values{"studentId": {"1001"}})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		} else {
			fmt.Println(string(body))
		}
	}
}

func TestGetGender(t *testing.T) {
	client := http.Client{}
	requestBody := map[string]interface{}{
		"studentId": "1001",
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Errorf("Failed to marshal JSON: %v", err)
		t.Fail()
		return
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:2345/get_gender", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		t.Fail()
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		t.Errorf("Failed to send request: %v", err)
		t.Fail()
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
		t.Fail()
		return
	}

	// 处理响应
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		t.Fail()
	} else {
		fmt.Println(string(body))
	}
}

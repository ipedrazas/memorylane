// main_test.go

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"),
		"localhost")

	ensureTableExists()

	code := m.Run()
	clearTable()
	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentTask(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/task/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Task not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Task not found'. Got '%s'", m["error"])
	}
}

func TestCreateTask(t *testing.T) {
	clearTable()

	payload := []byte(`{"name":"test task","completed":false}`)

	req, _ := http.NewRequest("POST", "/task", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test task" {
		t.Errorf("Expected task name to be 'test task'. Got '%v'", m["name"])
	}

	if m["completed"] == true {
		t.Errorf("Expected task to be not completed. Got '%v'", m["completed"])
	}

	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected task ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetTask(t *testing.T) {
	clearTable()
	addTasks(1)

	req, _ := http.NewRequest("GET", "/task/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateTask(t *testing.T) {
	clearTable()
	addTasks(1)

	req, _ := http.NewRequest("GET", "/task/1", nil)
	response := executeRequest(req)
	var originalTask map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalTask)

	payload := []byte(`{"name":"test task - updated name","completed":false}`)

	req, _ = http.NewRequest("PUT", "/task/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalTask["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalTask["id"], m["id"])
	}

	if m["name"] == originalTask["name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalTask["name"], m["name"], m["name"])
	}

	if m["completed"] == originalTask["completed"] {
		t.Errorf("Expected the completed to change from '%v' to '%v'. Got '%v'", originalTask["completed"], m["completed"], m["completed"])
	}
}

func TestDeleteTask(t *testing.T) {
	clearTable()
	addTasks(1)

	req, _ := http.NewRequest("GET", "/task/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/task/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/task/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM tasks")
	a.DB.Exec("ALTER SEQUENCE tasks_id_seq RESTART WITH 1")
}

func addTasks(count int) {
	if count < 1 {
		count = 1
	}

	flag := false
	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO tasks(name, completed) VALUES($1, $2)", "Task "+strconv.Itoa(i), !flag)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS tasks
(
    id SERIAL,
    name TEXT NOT NULL,
    completed BOOLEAN,
    CONSTRAINT tasks_pkey PRIMARY KEY (id)
)`

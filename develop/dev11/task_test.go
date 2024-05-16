package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Function for sending POST request and getting response
func sendPostRequest(handler http.HandlerFunc, url string, body string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

// Function for sending GET request and getting response
func sendGetRequest(handler http.HandlerFunc, url string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func TestCreateEvent(t *testing.T) {
	handler := http.HandlerFunc(createEventHandler)

	body := "user_id=1&date=2024-01-01&details=Test Event"
	rr := sendPostRequest(handler, "/create_event", body)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, rr.Code)
	}

	expected := `{"result":"Event created successfully"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected response body %s; got %s", expected, rr.Body.String())
	}
}

func TestUpdateEvent(t *testing.T) {
	handler := http.HandlerFunc(updateEventHandler)

	body := "user_id=1&date=2024-01-01&event_id=1&details=Updated Event"
	rr := sendPostRequest(handler, "/update_event", body)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, rr.Code)
	}

	expected := `{"result":"Event updated successfully"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected response body %s; got %s", expected, rr.Body.String())
	}
}

func TestDeleteEvent(t *testing.T) {
	handler := http.HandlerFunc(deleteEventHandler)

	body := "user_id=1&date=2024-01-01&event_id=1"
	rr := sendPostRequest(handler, "/delete_event", body)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, rr.Code)
	}

	expected := `{"result":"Event deleted successfully"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected response body %s; got %s", expected, rr.Body.String())
	}
}

func TestEventsForDay(t *testing.T) {

	handler := http.HandlerFunc(createEventHandler)
	body := "user_id=1&date=2024-01-01&details=Test Event"
	sendPostRequest(handler, "/delete_event", body)

	handler = http.HandlerFunc(eventsForDayHandler)
	url := "/events_for_day?user_id=1&date=2024-01-01"
	req := sendGetRequest(handler, url)

	if req.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, req.Code)
	}

	handler = http.HandlerFunc(deleteEventHandler)
	body = "user_id=1&date=2024-01-01&event_id=1"
	sendPostRequest(handler, "/delete_event", body)
}

func TestEventsForWeek(t *testing.T) {

	handler := http.HandlerFunc(createEventHandler)
	body := "user_id=1&date=2024-01-01&details=Test Event"
	sendPostRequest(handler, "/delete_event", body)
	body = "user_id=1&date=2024-01-02&details=Test Event 2"
	sendPostRequest(handler, "/delete_event", body)

	handler = http.HandlerFunc(eventsForDayHandler)
	url := "/events_for_day?user_id=1&date=2024-01-01"
	req := sendGetRequest(handler, url)

	if req.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, req.Code)
	}

	handler = http.HandlerFunc(deleteEventHandler)
	body = "user_id=1&date=2024-01-01&event_id=1"
	sendPostRequest(handler, "/delete_event", body)
	body = "user_id=1&date=2024-01-02&event_id=1"
	sendPostRequest(handler, "/delete_event", body)
}

func TestEventsForMonth(t *testing.T) {
	handler := http.HandlerFunc(createEventHandler)
	body := "user_id=1&date=2024-01-01&details=Test Event"
	sendPostRequest(handler, "/delete_event", body)
	body = "user_id=1&date=2024-01-25&details=Test Event 2"
	sendPostRequest(handler, "/delete_event", body)

	handler = http.HandlerFunc(eventsForDayHandler)
	url := "/events_for_day?user_id=1&date=2024-01-01"
	req := sendGetRequest(handler, url)

	if req.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, req.Code)
	}

	handler = http.HandlerFunc(deleteEventHandler)
	body = "user_id=1&date=2024-01-01&event_id=1"
	sendPostRequest(handler, "/delete_event", body)
	body = "user_id=1&date=2024-01-25&event_id=1"
	sendPostRequest(handler, "/delete_event", body)
}

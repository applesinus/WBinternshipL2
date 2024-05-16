package main

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503.
		В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400.
		В случае остальных ошибок сервер должен возвращать HTTP 500.
	4. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.

Код должен проходить проверки go vet и golint.
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Event - struct for event
type Event struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Date    string `json:"date"`
	Details string `json:"details"`
}

var mutex sync.Mutex

func checkUser(userID int) error {
	if _, err := os.Stat(fmt.Sprintf("user_%d", userID)); os.IsNotExist(err) {
		return fmt.Errorf("user %d does not exist", userID)
	}
	return nil
}

func checkEvent(userID int, date string, id int) error {
	folderName := fmt.Sprintf("user_%d", userID)
	fileName := fmt.Sprintf("%s/%s%v.txt", folderName, date, id)

	if err := checkUser(userID); err != nil {
		return err
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return fmt.Errorf("Event %s does not exist", date)
	}

	return nil
}

func createUser(userID int) error {
	folderName := fmt.Sprintf("user_%d", userID)
	mutex.Lock()
	defer mutex.Unlock()

	if err := checkUser(userID); err != nil {
		err := os.Mkdir(folderName, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func createEvent(userID int, date string, details string) error {
	folderName := fmt.Sprintf("user_%d", userID)

	if err := checkUser(userID); err != nil {
		createUser(userID)
	}

	files, err := os.ReadDir(folderName)

	if err != nil {
		return err
	}

	mutex.Lock()
	defer mutex.Unlock()

	max := -1
	for _, f := range files {
		if strings.Contains(f.Name(), date) {
			if temp, _ := strconv.Atoi(strings.Split(f.Name(), date)[1]); temp > max {
				max = temp
			}
		}
	}
	if max == -1 {
		max = 0
	}
	id := max + 1

	fileName := fmt.Sprintf("%s/%s%v.txt", folderName, date, id)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte(details))

	return nil
}

func updateEvent(userID int, date string, id int, details string) error {
	folderName := fmt.Sprintf("user_%d", userID)
	fileName := fmt.Sprintf("%s/%s%v.txt", folderName, date, id)

	mutex.Lock()
	defer mutex.Unlock()

	if err := checkUser(userID); err != nil {
		return err
	}

	if err := checkEvent(userID, date, id); err != nil {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Truncate(0)
	file.Write([]byte(details))
	return nil
}

func deleteEvent(userID int, date string, id int) error {
	folderName := fmt.Sprintf("user_%d", userID)
	fileName := fmt.Sprintf("%s/%s%v.txt", folderName, date, id)

	mutex.Lock()
	defer mutex.Unlock()

	if err := checkUser(userID); err != nil {
		return err
	}

	if err := checkEvent(userID, date, id); err != nil {
		return err
	}
	if err := os.Remove(fileName); err != nil {
		return err
	}

	files, err := os.ReadDir(folderName)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		if err := os.Remove(folderName); err != nil {
			return err
		}
	} else {
		for _, f := range files {
			if strings.Contains(f.Name(), date) {
				thisID, err := strconv.Atoi(strings.Split(f.Name(), date)[1])
				if err != nil {
					return err
				}
				if thisID > id {
					err := os.Rename(fileName, fmt.Sprintf("%s/%s%v.txt", folderName, date, thisID-1))
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func getEventsForDay(userID int, date string) ([]byte, error) {
	data := make([]byte, 0)
	folderName := fmt.Sprintf("user_%d", userID)

	mutex.Lock()
	defer mutex.Unlock()

	if err := checkUser(userID); err != nil {
		return nil, err
	}

	files, err := os.ReadDir(folderName)
	if err != nil {
		return nil, err
	}

	var file *os.File
	for _, f := range files {
		if strings.Contains(f.Name(), date) {
			file, err = os.Open(fmt.Sprintf("%s/%s", folderName, f.Name()))
			if err != nil {
				return nil, err
			}

			defer file.Close()
			fileData, err := os.ReadFile(fmt.Sprintf("%s/%s", folderName, f.Name()))
			if err != nil {
				return nil, err
			}
			data = append(fileData, '\n')
		}
	}

	return data, nil
}

func getEventsForWeek(userID int, date string) ([]byte, error) {
	data := make([]byte, 0)

	day, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 7; i++ {
		day = day.AddDate(0, 0, i)
		dayData, err := getEventsForDay(userID, day.Format("2006-01-02"))
		if err != nil {
			return nil, err
		}
		data = append(dayData, '\n')
	}

	return data, nil
}

func getEventsForMonth(userID int, date string) ([]byte, error) {
	data := make([]byte, 0)

	day, err := time.Parse("2006-01-02", date)
	month := day.Month()
	if err != nil {
		return nil, err
	}

	for i := 0; ; i++ {
		day = day.AddDate(0, 0, i)
		if day.Month() != month {
			break
		}
		dayData, err := getEventsForDay(userID, day.Format("2006-01-02"))
		if err != nil {
			return nil, err
		}
		data = append(dayData, '\n')
	}

	return data, nil
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parsing the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Getting the values from the form
	userID, err := strconv.Atoi(r.Form.Get("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	date := r.Form.Get("date")
	details := r.Form.Get("details")

	// Creating the event
	err = createEvent(userID, date, details)
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	// returning the success result
	response := map[string]string{"result": "Event created successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parsing the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Getting the values from the form
	userID, err := strconv.Atoi(r.Form.Get("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	date := r.Form.Get("date")
	eventID, err := strconv.Atoi(r.Form.Get("event_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	details := r.Form.Get("details")

	// Updating the event
	err = updateEvent(userID, date, eventID, details)
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	// Returning the success result
	response := map[string]string{"result": "Event updated successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parsing the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Getting the values from the form
	userID, err := strconv.Atoi(r.Form.Get("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	date := r.Form.Get("date")
	eventID, err := strconv.Atoi(r.Form.Get("event_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Deleting the event
	err = deleteEvent(userID, date, eventID)
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	// Returning the success result
	response := map[string]string{"result": "Event deleted successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parsing the form data
	queryParams := r.URL.Query()
	userID, err := strconv.Atoi(queryParams.Get("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	date := queryParams.Get("date")

	// Getting the events for the specified date
	events, err := getEventsForDay(userID, date)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	// Returning the success result
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(events)
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parsing the form data
	queryParams := r.URL.Query()
	userID, err := strconv.Atoi(queryParams.Get("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	date := queryParams.Get("date")

	// Getting the events for the specified week
	events, err := getEventsForWeek(userID, date)
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	// Returning the success result
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(events)
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parsing the form data
	queryParams := r.URL.Query()
	userID, err := strconv.Atoi(queryParams.Get("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	date := queryParams.Get("date")

	// Getting the events for the specified month
	events, err := getEventsForMonth(userID, date)
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	// Returning the success result
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(events)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	// HTTP handlers for buisness logic
	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)
	http.HandleFunc("/events_for_day", eventsForDayHandler)
	http.HandleFunc("/events_for_week", eventsForWeekHandler)
	http.HandleFunc("/events_for_month", eventsForMonthHandler)

	// HTTP handlers for static files
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)

	// Logging middleware
	http.Handle("/", loggingMiddleware(http.DefaultServeMux))

	// Starting the server
	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	htmlContent, err := os.ReadFile("check.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%s", htmlContent)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	faviconContent, err := os.ReadFile("favicon.ico")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/x-icon")
	w.Write(faviconContent)
}

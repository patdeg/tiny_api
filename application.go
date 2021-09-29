package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIParameter struct {
	A int `json:"a,omitempty"`
	B int `json:"b,omitempty"`
}

type APIResponse struct {
	Result       int    `json:"result,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

func GetBody(r *http.Request) ([]byte, error) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	if err != nil {
		log.Printf("Error while dumping request: %v", err)
		return []byte{}, err
	}
	return buffer.Bytes(), nil
}

func WriteJSON(w http.ResponseWriter, d interface{}) error {
	jsonData, err := json.Marshal(d)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", jsonData)
	return nil
}

func ReadJSON(b []byte, d interface{}) error {
	return json.Unmarshal(b, d)
}

func MyFunction(param APIParameter, response *APIResponse) error {
	response.Result = param.A + param.B
	return nil
}

// API handler
// Expect body with JSON such as {'a':2, 'b':2}
// Respond with 4
func APIHandler(w http.ResponseWriter, r *http.Request) {

	var response APIResponse

	body, err := GetBody(r)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		response.ErrorMessage = err.Error()
		WriteJSON(w, &response)
		return
	}

	var param APIParameter
	if err = json.Unmarshal(body, &param); err != nil {
		log.Printf("Body: %s", body)
		log.Printf("Error reading JSON from body: %v", err)
		response.ErrorMessage = err.Error()
		WriteJSON(w, &response)
		return
	}

	if err = MyFunction(param, &response); err != nil {
		log.Printf("Body: %s", body)
		log.Printf("Error running function: %v", err)
		response.ErrorMessage = err.Error()
		WriteJSON(w, &response)
		return
	}

	WriteJSON(w, &response)
}

func main() {

	// Define HTTP Router
	r := mux.NewRouter()
	r.HandleFunc("/", APIHandler)
	http.Handle("/", r)

	PORT := "8080"
	// Start application - port 8080 within the Docker image,
	log.Printf("Starting up on %v", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))

}

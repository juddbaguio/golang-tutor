package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type SampleHandler struct{}

func (sh *SampleHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world")) // sending data to client
}

func (sh *SampleHandler) WithStatusCode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated) // setting status code
	w.Write([]byte("hello world"))
}

func (sh *SampleHandler) WithContentType(w http.ResponseWriter, r *http.Request) {
	simpleMap := map[string]interface{}{
		"fullName": "Jim Xel Maghanoy",
		"age":      25,
	}

	jsonByte, err := json.Marshal(&simpleMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}

type PersonWithoutJSONTag struct {
	FullName     string
	Age          int
	AnotherField bool
}

type PersonWithJSONTag struct {
	FullName     string `json:"fullName"`
	Age          int    `json:"age"`
	AnotherField bool   `json:"-"` // this json tag will ignore when marshalling the field
}

func (sh *SampleHandler) WithJSONResponse(w http.ResponseWriter, r *http.Request) {
	// person := PersonWithoutJSONTag{
	// 	FullName:     "Jim Xel Maghanoy",
	// 	Age:          25,
	// 	AnotherField: true,
	// }

	person := PersonWithJSONTag{
		FullName:     "Jim Xel Maghanoy",
		Age:          25,
		AnotherField: true,
	}

	jsonByte, err := json.Marshal(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}

func main() {
	var sampleHandler *SampleHandler = &SampleHandler{}

	router := chi.NewRouter()

	router.Get("/", sampleHandler.HelloWorld)
	router.Get("/status-code", sampleHandler.WithStatusCode)
	router.Get("/content-type", sampleHandler.WithContentType)
	router.Get("/json", sampleHandler.WithJSONResponse)

	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

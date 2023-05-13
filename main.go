package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type SampleHandler struct{}

func (sh *SampleHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world")) // sending data to client
}

// with URL Parameters
func (sh *SampleHandler) HelloWorldUpgraded(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	w.Write([]byte("hello world: " + name)) // sending data to client
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

func (sh *SampleHandler) WithJSONBodyPayload(w http.ResponseWriter, r *http.Request) {
	var payload PersonWithJSONTag
	err := json.NewDecoder(r.Body).Decode(&payload)

	jsonResponseEncoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonResponseEncoder.Encode(map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponseEncoder.Encode(payload) // direct to the point, sending it back to the client
}

func (sh *SampleHandler) WithQueryParameters(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	log.Println(queryParams.Has("age"))

	// decima: 0,1,...9,10,11,
	// hexadecimal: 0,1,2,3,4,5,10,11,12,13,14,15,20,21
	// strconv.ParseBool for boolean
	// strconv.ParseBool()
	age, err := strconv.ParseInt(queryParams.Get("age"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("age is not a number"))
		return
	}

	log.Println(age)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(nil) // direct to the point, sending it back to the client
}

func main() {
	var sampleHandler *SampleHandler = &SampleHandler{}

	router := chi.NewRouter()

	// nouns/adjective
	// "/set-active-user" (very wrong)
	// "/delete-user" (very wrong)
	// [PUT] "/users/status"
	// [DELETE] "/users"
	// [DELETE] "/users/{userId}"

	router.Get("/", sampleHandler.HelloWorld)
	router.Get("/hello-world/{name}", sampleHandler.HelloWorldUpgraded)
	router.Get("/status-code", sampleHandler.WithStatusCode)
	router.Get("/content-type", sampleHandler.WithContentType)
	router.Get("/query-params", sampleHandler.WithQueryParameters)
	router.Get("/json", sampleHandler.WithJSONResponse)
	router.Post("/json-body", sampleHandler.WithJSONBodyPayload) // or PUT, DELETE, PATCH

	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

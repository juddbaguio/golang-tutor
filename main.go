package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

func (sh *SampleHandler) WithFileUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile") // file <--- [][][][] <----- [][][][]
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close() // closes reader connection, expects no data to be received from the client
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := os.Create("./temp-images/" + handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	// fileByte -> actual content
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func main() {
	var sampleHandler *SampleHandler = &SampleHandler{}

	// creates temporary-directory
	os.Mkdir("./temp-images", os.ModeAppend.Perm())

	router := chi.NewRouter()

	// nouns/adjective
	// "/set-active-user" (very wrong)
	// "/delete-user" (very wrong)
	// [PUT/PATCH] "/users/status"
	// [DELETE] "/users"
	// [DELETE] "/users/{userId}"

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/", sampleHandler.HelloWorld)
	router.Get("/hello-world/{name}", sampleHandler.HelloWorldUpgraded)
	router.Get("/status-code", sampleHandler.WithStatusCode)
	router.Get("/content-type", sampleHandler.WithContentType)
	router.Get("/query-params", sampleHandler.WithQueryParameters)
	router.Get("/json", sampleHandler.WithJSONResponse)
	router.Post("/file-upload", sampleHandler.WithFileUpload)
	router.Put("/file-upload", sampleHandler.WithFileUpload)
	router.Post("/json-body", sampleHandler.WithJSONBodyPayload) // or PUT, DELETE, PATCH
	router.Put("/json-body", sampleHandler.WithJSONBodyPayload)
	router.Delete("/json-body", sampleHandler.WithJSONBodyPayload)
	router.Patch("/json-body", sampleHandler.WithJSONBodyPayload)
	router.Get("/websocket", sampleHandler.WebsocketHandler)

	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Request struct (
	ID float64 `json:"id"`
	Value String `json:"value"`
)

type Response struct (
	Message string `json:"message"`
	OK bool `json:"ok"`
)


const (
	STATIC_PATH = "/static/"
	PORT       = "8080"
	ENV		= ".env"
	UPLOAD_PATH	= "/uploadf"
	LANDING_PATH = "/replace"
	STATIC_CONTENT_ENV_VAR = "static_content"
	POST_FORM_NAME_FIELD = "rname"
	POST_FORM_FILE_FIELD = "myFile"
	REPLACEMENT_STRING = "Unnamed"
)

var myClient = &http.Client{}

func main() {
	lambda.Start(Handler)
}

func Handler(request Request)(Reponse, error){

	return Response{
		
	}
	fmt.Println("Starting server....")

	NewRouter()

}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(ENV)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func NewRouter() {
	fmt.Println("Configuring router...")

	static_file := goDotEnvVariable(STATIC_CONTENT_ENV_VAR)
	
	fmt.Println("dir..", static_file)

	r := mux.NewRouter()

	r.HandleFunc(UPLOAD_PATH, uploadFile)

	r.PathPrefix(LANDING_PATH).Handler( http.StripPrefix(LANDING_PATH, http.FileServer(http.Dir(static_file))))

	srv := http.Server{
		Handler: r,
		Addr:    ":"+PORT,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uploading file..")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	replaceString := r.FormValue(POST_FORM_NAME_FIELD)

	fmt.Println("Remove name..", replaceString)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile(POST_FORM_FILE_FIELD)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-files", "upload-*.html")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	updatedString := strings.ReplaceAll(string(fileBytes), replaceString, REPLACEMENT_STRING)
	// write this byte array to our temporary file
	tempFile.Write([]byte(updatedString))
	// return that we have successfully uploaded our file!
	fmt.Println("Printing file to streamWriter.")
	fmt.Fprint(w, updatedString)
	// fmt.Fprintf(w, "Successfully Uploaded File\n")
}
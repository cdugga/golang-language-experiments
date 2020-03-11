package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	STATIC_DIR = "../static/"
	STATIC_PATH = "/static/"
	PORT       = "8080"
)

var myClient = &http.Client{}

func main(){
	log.Println("Starting server....")
	r := NewRouter()

	r.HandleFunc("/uploadf", uploadFile)
	err := http.ListenAndServe(":"+PORT, r) //, GreetingHandler("Mayo 4 GO"))
	if err != nil{
		log.Fatal(err)
	}
}

func NewRouter() *mux.Router {
	log.Println("Configuring router...")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dir..", dir)
	r := mux.NewRouter()

	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	return r
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uploading file..")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	replaceString := r.FormValue("rname")

	fmt.Println("Remove name..", replaceString)


	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
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

	updatedString := strings.ReplaceAll(string(fileBytes), replaceString, "Unnamed")
	// write this byte array to our temporary file
	tempFile.Write([]byte(updatedString))
	// return that we have successfully uploaded our file!
	fmt.Println("Printing file to streamWriter.")
	fmt.Fprint(w, updatedString)
	// fmt.Fprintf(w, "Successfully Uploaded File\n")
}
package main

import (
	"net/http"
	"html/template"
)

type Bucket struct {
	BucketName string
	CreationDate string
}

type Lambda struct {
	FunctionName string
	FunctionArn string
}

type DynamoDB struct  {
	DatabaseName string
}

type ApiGateway struct {
	Name string
	Description string
}

func (a *AWSCloud) ShowHandler(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("inventory.html"))
	tmpl.ExecuteTemplate(w, "inventory.html", a)
}

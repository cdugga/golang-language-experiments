
# HTMLReplace Golang Utility


![GitHub contributors](https://github.com/cdugga/golang-language-experiments/blob/master/html_stream_replace/Readme.md)

HTMLReplace is a GO utility application which uses native string operations to replace occurences of a given string in a HTLM document. Two parameters are supplied by the end user [1.the HTML file and 2.the string to replace]. The utility was created specifically for masking player names which appear in the automatically generated results of a bridgemate device. 

## Installation

The application is built as a [go modules](https://blog.golang.org/using-go-modules).

Install dependencies
```go
go mod init main.go
```

Build and Run localy
```go
go build -o htmlreplace
./htmlreplace
```

## Dockerize

Docker build and run
```bash
docker build -t htmlreplace:1.1 .
###Run attached
docker run -a stdin -a stdout -p 8080:8080 htmlreplace:1.1
###Run detached
docker run -d -p 8081:8080 htmlreplace:1.1
###Run with shell
docker run -it -p 3002:3002 htmlreplace:1.1 /bin/bash
```

## Deploymeny
```bash
### build image to container registry
gcloud builds submit --tag gcr.io/cloudrun-sample-270320/htmlreplace

### deploy to cloud run
gcloud run deploy --image gcr.io/cloudrun-sample-270320/htmlreplace --platform managed
```

##Usage

```bash

```





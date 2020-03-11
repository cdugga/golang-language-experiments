
#Initialize a go.mod to declare the module
go mod init main.go

#running locally
go build -o apiserver
./apiserver

#docker util commands
##container pid
docker inspect -f '{{.State.Pid}}' ac6388994f47
##netstat
sudo nsenter -t 15652 -n netstat

#build image
docker build -t apiserver:1.1 .

#run detached
docker run -d -p 8081:8080 apiserver:1.1

#run attached
docker run -a stdin -a stdout -p 3002:3002 apiserver:1.1

#run with shell
docker run -it -p 3002:3002 apiserver:1.1 /bin/bash

# build image to container registry
gcloud builds submit --tag gcr.io/cloudrun-sample-270320/htmlreplace

# deploy to cloud run
gcloud run deploy --image gcr.io/cloudrun-sample-270320/htmlreplace --platform managed
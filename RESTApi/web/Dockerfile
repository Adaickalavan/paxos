# Start from a Debian image with the latest version of GO 
# installed and a workspace (GOPATH) configure at /go.
FROM golang:latest

#Author label
LABEL Author Adaickalavan Meiyappan

#Copy the local package files (from development computer) to the container's workspace (docker image)
COPY . /go/src/app 

# BUild the app command inside the container
RUN go install app

# Document port which the service listens on
EXPOSE 8080

#Run the app command by default when the container starts
ENTRYPOINT /go/bin/app

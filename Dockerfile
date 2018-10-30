# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

RUN mkdir /go/src/app

# install golang dependency management tool
RUN go get -u github.com/golang/dep/cmd/dep


# Add main file to /go/src/app
ADD ./main.go /go/src/app

# Copy Gopkg files to install dependencies
COPY ./Gopkg.toml /go/src/app
COPY ./Gopkg.lock /go/src/app

# Include xlsx file into app
COPY ./dummy.csv /go/src/app

WORKDIR /go/src/app 

RUN dep ensure 
RUN go test -v 
RUN go build

CMD ["./app"]
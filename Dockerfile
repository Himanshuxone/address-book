# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

RUN mkdir /go/src/app

# install golang dependency management tool
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/app 

# Add main file to /go/src/app
ADD ./main.go /go/src/app
COPY ./parsecsv /go/src/app/parsecsv

# Copy Gopkg files to install dependencies
COPY ./Gopkg.toml /go/src/app
COPY ./Gopkg.lock /go/src/app

# Include xlsx file into app
COPY ./dummy.csv /go/src/app

RUN dep ensure -update -v
RUN go test -v 
RUN go build

CMD ["./app"]